package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/dpcat237/afs/internal/interfaces"
	"github.com/dpcat237/afs/internal/model"
)

const (
	serverErrorMessage = `{ "message": "Server error" }`
	successfulResponse = `{ "message": "Success" }`
)

type controller struct {
	lgr interfaces.Logger
}

func newController(lgr interfaces.Logger) *controller {
	return &controller{lgr: lgr}
}

//GetBodyContent gets the body content and converts to provided struct.
func (cnt controller) GetBodyContent(req *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, req.ContentLength))
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

//HealthCheck writes successful response.
func (cnt controller) HealthCheck(w http.ResponseWriter, req *http.Request) {
	if _, err := w.Write([]byte(successfulResponse)); err != nil {
		cnt.lgr.Warn("Error returning health check " + err.Error())
	}
}

//ReturnError writes error message in response.
func (cnt controller) ReturnError(w http.ResponseWriter, out model.Output) {
	w.WriteHeader(out.Status)
	if err := json.NewEncoder(w).Encode(out); err != nil {
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}
}

//ReturnJson converts struct to JSON and writes it to the response.
func (cnt controller) ReturnJson(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}
}
