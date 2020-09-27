package model

import (
	"errors"
	"net/http"
)

// Output contains error details.
type Output struct {
	Error   error  `json:"-"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

//NewErrorNil creates empty Output.
func NewErrorNil() Output {
	return Output{}
}

//Create creates Output object from error message.
func NewErrorString(msg string) Output {
	return Output{
		Error:   errors.New(msg),
		Message: msg,
		Status:  http.StatusInternalServerError,
	}
}

//MessageLog returns message for logging.
func (out Output) MessageLog() string {
	if out.Message == out.Error.Error() {
		return out.Message
	}
	return out.Message + ": " + out.Error.Error()
}

//IsError checks if Output contains error.
func (out Output) IsError() bool {
	return out.Error != nil
}

//WithError adds an error and returns Output.
func (out Output) WithError(err error) Output {
	if err == nil && out.Error != nil {
		return out
	}
	out.Error = err
	return out
}

//WithStatus adds error's status and returns Output.
func (out Output) WithStatus(status int) Output {
	out.Status = status
	return out
}
