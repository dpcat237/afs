package handler

import (
	"context"
	"net/http"

	"github.com/dpcat237/afs/internal/model"
)

type link struct {
}

func newLink() *link {
	return &link{}
}

//ProcessLinks after data validation process requested links and returns responses.
func (hnd link) ProcessLinks(ctx context.Context, lnkReq model.LinksRequest) (model.LinksResponse, model.Output) {
	var rsp model.LinksResponse
	if lnkReq.ExceedsLimit() {
		return rsp, model.NewErrorString("Exceeds limit of links").WithStatus(http.StatusPreconditionFailed)
	}
	if lnkReq.IsEmpty() {
		return rsp, model.NewErrorString("No links to process").WithStatus(http.StatusPreconditionFailed)
	}

	pl := model.NewPool(ctx, lnkReq)
	return pl.MakeWork()
}
