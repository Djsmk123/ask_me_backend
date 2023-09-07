package auth

import (
	"net/http"
	"strings"

	errorhandler "github.com/djsmk123/askmeapi/api/error_handler"
	"github.com/djsmk123/askmeapi/utils"
)

func (a *AuthUtils) Login(email string, password string, fcmToken string) (*UserResponseTypeWithToken, *errorhandler.ErrorHandlerApi) {

	user, err := a.Database.GetUserByEmail(a.ctx, email)
	if (err) != nil {
		return nil, UserErrorHandler(err)
	}
	if strings.ToLower(user.Provider) != "password" {
		return nil, &errorhandler.ErrorHandlerApi{
			Error:      ErrAnonymousUserFound,
			StatusCode: http.StatusForbidden,
		}
	}
	err = utils.CheckPassword(password, user.PasswordHash.String)
	if err != nil {
		return nil, UserErrorHandler(err)
	}

	rsp, e := a.CreateUserObjectForAuth(user, fcmToken)
	if e != nil {
		return nil, e
	}
	return rsp, nil

}
