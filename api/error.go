package api

import "errors"

var (
	errUserAlreadyExist = errors.New("user exist already")
	errWrongPassword    = errors.New("please check credentials")
	errUserNotExist     = errors.New("user not exist")

	errMissingAuthHeader    = errors.New("missing authorization header")
	errInvalidAuthHeader    = errors.New("invalid authorization header")
	errVerifiyingAuthHeader = errors.New("invalid authorization key")
)
