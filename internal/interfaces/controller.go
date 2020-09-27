package interfaces

import (
	"net/http"

	"github.com/dpcat237/afs/internal/model"
)

//Controller contains methods for HTTP controllers.
type Controller interface {
	GetBodyContent(req *http.Request, data interface{}) error
	HealthCheck(w http.ResponseWriter, req *http.Request)
	ReturnError(w http.ResponseWriter, out model.Output)
	ReturnJson(w http.ResponseWriter, v interface{})
}
