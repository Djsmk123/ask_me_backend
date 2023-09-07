package auth

import (
	"database/sql"
	"errors"
	"net/http"

	errorhandler "github.com/djsmk123/askmeapi/api/error_handler"
	"github.com/lib/pq"
)

var (
	ErrUserAlreadyExist   = errors.New("user exist already")
	ErrWrongPassword      = errors.New("invalid credentials")
	ErrUserNotExist       = errors.New("user not exist")
	ErrAnonymousUserFound = errors.New("not authenticated to requesting this")
)

func UserErrorHandler(err error) *errorhandler.ErrorHandlerApi {
	if err == sql.ErrNoRows {
		return &errorhandler.ErrorHandlerApi{
			Error:      ErrUserNotExist,
			StatusCode: http.StatusNotFound,
		}
	}
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation", "unique_key_voilation":
			return &errorhandler.ErrorHandlerApi{
				Error:      ErrUserAlreadyExist,
				StatusCode: http.StatusForbidden,
			}
		}
	}
	return errorhandler.InternalServerErrorHandler(err)

}
