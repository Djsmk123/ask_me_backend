package auth

import errorhandler "github.com/djsmk123/askmeapi/api/error_handler"

func (a *AuthUtils) GetUser(userId int32) (*UserResponseType, *errorhandler.ErrorHandlerApi) {
	user, err := a.Database.GetUserByID(a.ctx, userId)

	if (err) != nil {
		return nil, UserErrorHandler(err)
	}
	rps := a.CreateUserObjectWithoutToken(user)
	return &rps, nil
}
