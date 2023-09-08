package auth

import (
	errorhandler "github.com/djsmk123/askmeapi/api/error_handler"
)

type DeleteUserResponse struct {
	message string
}

func (a *AuthUtils) DeleteUser(userId int32) (*DeleteUserResponse, *errorhandler.ErrorHandlerApi) {
	//delete question where user_id =id
	a.Database.DeleteAnswerByUserId(a.ctx, userId)

	a.Database.DeleteQuestionByUserId(a.ctx, userId)

	a.Database.DeleteFcmTokenByUserId(a.ctx, userId)
	a.Database.DeleteJWTokenByUserId(a.ctx, userId)

	_, err := a.Database.DeleteUserById(a.ctx, userId)
	if (err) != nil {
		return nil, UserErrorHandler(err)
	}
	return &DeleteUserResponse{
		message: "deleted-succesfully",
	}, nil

}
