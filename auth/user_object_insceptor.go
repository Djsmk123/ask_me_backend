package auth

import (
	"time"

	errorhandler "github.com/djsmk123/askmeapi/api/error_handler"
	db "github.com/djsmk123/askmeapi/db/sqlc"
)

type UserResponseTypeWithToken struct {
	AccessToken string           `json:"token"`
	User        UserResponseType `json:"user"`
}

type UserResponseType struct {
	ID                  int32     `json:"id"`
	Username            string    `json:"username"`
	Email               string    `json:"email"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	PublicProfileImage  string    `json:"public_profile_image"`
	PrivateProfileImage string    `json:"private_profile_image"`
}

func (a *AuthUtils) CreateUserObjectForAuth(user db.User, fcmToken string) (*UserResponseTypeWithToken, *errorhandler.ErrorHandlerApi) {

	t, err := a.CreateJwtToken(int64(user.ID), user.Username)
	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	authpayload, err := a.TokenMaker.VerifyToken(t)
	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)

	}
	rsp := UserResponseTypeWithToken{
		AccessToken: t,
		User:        GetUserResponse(user),
	}
	err = a.CreateFcmToken(fcmToken, user, int(authpayload.ID), a.Database)
	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	return &rsp, nil

}
func (a *AuthUtils) CreateUserObjectWithoutToken(user db.User) UserResponseType {
	return GetUserResponse(user)
}

func GetUserResponse(user db.User) UserResponseType {
	return UserResponseType{
		ID:                  user.ID,
		Username:            user.Email,
		Email:               user.Email,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           user.UpdatedAt,
		PublicProfileImage:  user.PublicProfileImage,
		PrivateProfileImage: user.PrivateProfileImage,
	}
}
