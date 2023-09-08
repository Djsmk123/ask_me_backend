package errorhandler

import (
	"errors"
	"net/http"
)

var (
	ErrMissingAuthHeader    = errors.New("missing authorization header")
	ErrInvalidAuthHeader    = errors.New("invalid authorization header")
	ErrVerifiyingAuthHeader = errors.New("invalid authorization key")
)

type ErrorHandlerApi struct {
	Error      error
	StatusCode int
}

func InternalServerErrorHandler(error error) *ErrorHandlerApi {
	return &ErrorHandlerApi{
		Error:      error,
		StatusCode: http.StatusInternalServerError,
	}
}
