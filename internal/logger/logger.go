package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// logger wraps zap.Logger
type logger struct {
}

type requestLog struct {
	Action       string `json:"action"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
	CreatedAt    int64  `json:"created_at"`
	ResponseTime string `json:"response_time"`
}

//Fatal logs a message at fatal level then calls os.Exit(1).
func (lgr logger) Fatal(msg string) {
	log.Fatal(msg)
}

//Info logs a message at info level.
func (lgr logger) Info(msg string) {
	log.Println(msg)
}

//Infof uses fmt.Sprintf to log a message at info level.
func (lgr logger) Infof(template string, args ...interface{}) {
	log.Printf(template, args...)
}

//Infow log a message at info level with additional information as key-value pairs.
func (lgr logger) Infow(msg string, keysAndValues ...interface{}) {
	log.Println(msg, keysAndValues)
}

// New initializes a new logger.
func New() (*logger, error) {
	return &logger{}, nil
}

// RequestEnd log end of HTTP request process.
func (lgr logger) RequestEnd(act string, startAt time.Time, status *int, errMsg *string) {
	req := requestLog{
		Action:       act,
		StatusCode:   *status,
		ErrorMessage: *errMsg,
		CreatedAt:    startAt.Unix(),
		ResponseTime: fmt.Sprintf("%.4f", time.Since(startAt).Seconds()),
	}
	txt, err := json.Marshal(req)
	if err != nil {
		lgr.Info("Error marshaling request log")
		return
	}

	log.Println("http_request", string(txt))
}

//Warn logs a message at warning level.
func (lgr logger) Warn(msg string) {
	log.Println("Warning:", msg)
}
