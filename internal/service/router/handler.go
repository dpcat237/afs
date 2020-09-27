package router

import (
	"net/http"
	"strings"
)

type httpHandler struct {
	endpoint *Endpoint
}

func newHandler(endpoint *Endpoint) *httpHandler {
	return &httpHandler{
		endpoint: endpoint,
	}
}

func (hnd *httpHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Serve endpoint if not on root one
	if hnd.endpoint != hnd.endpoint.root {
		hnd.endpoint.serve(rw, r)
		return
	}

	path := r.URL.Path
	if path != "/" && strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	for _, e := range hnd.endpoint.root.endpoints {
		ctx, ok := e.pattern.match(r.Context(), path)
		if ok {
			e.chain.ServeHTTP(rw, r.WithContext(ctx))
			return
		}
	}

	rw.WriteHeader(http.StatusNotFound)
}
