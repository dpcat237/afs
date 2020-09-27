package interfaces

import (
	"context"
	"net/http"

	"github.com/dpcat237/afs/internal/model"
)

//LinkController contains methods to process HTTP requests about Link.
type LinkController interface {
	ProcessLinks(w http.ResponseWriter, r *http.Request)
}

//LinkHandler contains methods related to Link processes.
type LinkHandler interface {
	ProcessLinks(ctx context.Context, lnkReq model.LinksRequest) (model.LinksResponse, model.Output)
}
