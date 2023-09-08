package auth

import (
	"database/sql"

	errorhandler "github.com/djsmk123/askmeapi/api/error_handler"
	db "github.com/djsmk123/askmeapi/db/sqlc"
)

type LogoutUserResponse struct {
	message string
}

func (a *AuthUtils) Logout(userId int32, token string) (*LogoutUserResponse, *errorhandler.ErrorHandlerApi) {

	arg := db.UpdateJwtTokenParams{
		JwtToken: token,
		UserID:   userId,
	}
	_, err := a.Database.UpdateJwtToken(a.ctx, arg)

	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	//update fcm toke if exist
	arg1 := db.UpdateFcmTokenParams{
		ID:      userId,
		IsValid: false,
	}
	_, err = a.Database.UpdateFcmToken(a.ctx, arg1)

	if err != nil && err != sql.ErrNoRows {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	return &LogoutUserResponse{
		message: "Logout user successfully",
	}, nil
}
