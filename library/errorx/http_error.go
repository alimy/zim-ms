package errorx

import "net/http"

type httpError struct {
	code int
	msg  string
}

func (e *httpError) Error() string {
	return e.msg
}

func (e *httpError) Code() int {
	return e.code
}

// NewHttpError get a http error
func NewHttpError(code int, msg string) HttpError {
	return &httpError{
		code: code,
		msg:  msg,
	}
}

// InternalServerError get an internal server error
func InternalServerError(err error) HttpError {
	return &httpError{
		code: http.StatusInternalServerError,
		msg:  err.Error(),
	}
}
