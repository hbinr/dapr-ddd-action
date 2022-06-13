package errorx

import (
	"fmt"
	"net/http"
)

type Error struct {
	Type     string   `json:"type"`
	Code     int      `json:"code"`
	Message  string   `json:"message"`
	Metadata Metadata `json:"-"`
	Err      error    `json:"-"`
}

type Metadata map[string]interface{}

func Internal(err error, format string, args ...interface{}) *Error {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, message).
		WithMetadata(Metadata{"error": err}).
		WithError(err)
}

func NotFound(format string, args ...interface{}) *Error {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusNotFound), http.StatusNotFound, message)
}

func Unauthorized(format string, args ...interface{}) *Error {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, message)
}

func BadRequest(format string, args ...interface{}) *Error {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, message)
}

func From(err error) *Error {
	if err == nil {
		return nil
	}
	if errx, ok := err.(*Error); ok {
		return errx
	}
	return Internal(err, err.Error())
}

func New(t string, code int, message string) *Error {
	return &Error{
		Type:    t,
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) WithMessage(format string, args ...interface{}) *Error {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

func (e *Error) WithMetadata(metadata Metadata) *Error {
	e.Metadata = metadata
	return e
}

func (e *Error) WithError(err error) *Error {
	e.Err = err
	return e
}
