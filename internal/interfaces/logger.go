package interfaces

import "time"

//Logger contains methods for logging.
type Logger interface {
	Fatal(msg string)
	Info(msg string)
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	RequestEnd(act string, startAt time.Time, status *int, errMsg *string)
	Warn(msg string)
}
