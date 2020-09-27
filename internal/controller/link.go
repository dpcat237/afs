package controller

import (
	"net/http"
	"time"

	"github.com/dpcat237/afs/internal/interfaces"
	"github.com/dpcat237/afs/internal/model"
)

type link struct {
	cnt    interfaces.Controller
	lgr    interfaces.Logger
	lnkHnd interfaces.LinkHandler
}

func newLink(cnt interfaces.Controller, lgr interfaces.Logger, lnkHnd interfaces.LinkHandler) *link {
	return &link{
		cnt:    cnt,
		lgr:    lgr,
		lnkHnd: lnkHnd,
	}
}

//ProcessLinks reads requested links and returns responses.
func (ctr link) ProcessLinks(w http.ResponseWriter, req *http.Request) {
	status := http.StatusOK
	errMsg := ""
	defer ctr.lgr.RequestEnd("link.process", time.Now(), &status, &errMsg)

	var lnksReq model.LinksRequest
	if err := ctr.cnt.GetBodyContent(req, &lnksReq); err != nil {
		status = http.StatusPreconditionFailed
		errMsg = err.Error()
		ctr.cnt.ReturnError(w, model.NewErrorString(errMsg).WithStatus(status))
		return
	}

	rsp, out := ctr.lnkHnd.ProcessLinks(req.Context(), lnksReq)
	if out.IsError() {
		status = out.Status
		errMsg = out.MessageLog()
		ctr.cnt.ReturnError(w, model.NewErrorString(out.Message).WithStatus(status))
		return
	}

	w.WriteHeader(status)
	ctr.cnt.ReturnJson(w, rsp)
}
