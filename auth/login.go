package auth

import (
	"database/sql"
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
		return nil, &errorhandler.ErrorHandlerApi{
			StatusCode: http.StatusForbidden,
			Error:      ErrWrongPassword,
		}
	}

	rsp, e := a.CreateUserObjectForAuth(user, fcmToken)
	if e != nil {
		return nil, e
	}
	return rsp, nil

}

func (a *AuthUtils) SocialLogin(email string, provider string, fcmToken string, image string) (*UserResponseTypeWithToken, *errorhandler.ErrorHandlerApi) {
	user, err := a.Database.GetUserByEmail(a.ctx, email)
	if (err) != nil {
		if err == sql.ErrNoRows {
			// create a new user
			return a.CreateUserWithProvider(email, provider, fcmToken, image)
		}
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	// check if user is already exist as anonymous user or any another provider
	if !strings.EqualFold(provider, user.Provider) {
		return nil, &errorhandler.ErrorHandlerApi{
			Error:      ErrAnonymousUserFound,
			StatusCode: http.StatusBadRequest,
		}
	}
	rsp, e := a.CreateUserObjectForAuth(user, fcmToken)
	if e != nil {
		return nil, e
	}
	return rsp, nil

}
