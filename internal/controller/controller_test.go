package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestLogger struct{}

func (TestLogger) Fatal(msg string) {}
func (TestLogger) Info(msg string) {
}
func (TestLogger) Infof(template string, args ...interface{}) {
}
func (TestLogger) Infow(msg string, keysAndValues ...interface{}) {
}
func (TestLogger) RequestEnd(act string, startAt time.Time, status *int, errMsg *string) {
}
func (TestLogger) Warn(msg string) {
}

func TestController_healthCheck(t *testing.T) {
	cnt := newController(TestLogger{})

	req, _ := http.NewRequest("GET", "/services/health", nil)
	w := httptest.NewRecorder()
	cnt.HealthCheck(w, req)
	if http.StatusOK != w.Code {
		t.Error("Wrong health response with status code", w.Code)
	}
}
