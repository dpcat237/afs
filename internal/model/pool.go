package model

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	requestTimeout = 1 * time.Second
	workersLimit   = 4
)

//Pool contains workers pool data.
type Pool struct {
	context   context.Context
	links     chan string
	output    Output
	responses chan LinkResponse
}

//NewPool creates workers pool with defined data.
func NewPool(ctx context.Context, lnkReq LinksRequest) Pool {
	var pl Pool
	size := len(lnkReq)

	pl.context = ctx
	pl.links = make(chan string, size)
	pl.responses = make(chan LinkResponse, size)

	for _, lnk := range lnkReq {
		pl.links <- lnk
	}
	close(pl.links)

	return pl
}

//MakeWork process requests concurrently with context or response error cancellation.
func (pl *Pool) MakeWork() (LinksResponse, Output) {
	var wg sync.WaitGroup
	for i := 0; i < workersLimit; i++ {
		wg.Add(1)
		go pl.worker(&wg)
	}
	wg.Wait()
	close(pl.responses)

	return pl.responseToSlice(), pl.output
}

func (pl Pool) processRequest(ctx context.Context, lnk string) (LinkResponse, Output) {
	var rsp LinkResponse
	req, err := http.NewRequest(http.MethodGet, lnk, nil)
	if err != nil {
		return rsp, NewErrorString(fmt.Sprintf("Error creating request for URL %s", lnk))
	}

	cli := http.Client{
		Timeout: requestTimeout,
	}

	cliRsp, err := cli.Do(req.WithContext(ctx))
	if err != nil {
		return rsp, NewErrorString("Error to URL " + lnk).WithStatus(http.StatusBadRequest).WithError(err)
	}
	if cliRsp.StatusCode >= http.StatusBadRequest {
		return rsp, NewErrorString(fmt.Sprintf("Error with status %d to URL %s", cliRsp.StatusCode, lnk))
	}

	rsp.URL = lnk
	rsp.Status = cliRsp.StatusCode
	return rsp, NewErrorNil()
}

func (pl Pool) responseToSlice() []LinkResponse {
	var rspSlc []LinkResponse
	if pl.output.IsError() {
		return rspSlc
	}

	for rsp := range pl.responses {
		rspSlc = append(rspSlc, rsp)
	}
	return rspSlc
}

func (pl *Pool) worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for lnk := range pl.links {
		if pl.context.Err() != nil || pl.output.IsError() {
			return
		}

		rsp, out := pl.processRequest(pl.context, lnk)
		if out.IsError() {
			pl.output = out
			return
		}
		pl.responses <- rsp
	}
}
