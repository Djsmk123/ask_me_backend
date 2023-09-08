package auth

import (
	"database/sql"
	"fmt"

	errorhandler "github.com/djsmk123/askmeapi/api/error_handler"
	db "github.com/djsmk123/askmeapi/db/sqlc"
	"github.com/djsmk123/askmeapi/utils"
)

func (a *AuthUtils) CreateUser(email string, password string, fcmToken string) (*UserResponseTypeWithToken, *errorhandler.ErrorHandlerApi) {
	//check if username exist or not
	hashPassword, err := utils.HashPassword(password)

	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	username := utils.RandomUserName(email)
	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	publicProviderImage := utils.RandomUserProfileImage()

	arg := db.CreateUserParams{
		Username:            username,
		PasswordHash:        sql.NullString{String: hashPassword, Valid: true},
		Email:               email,
		Provider:            "password",
		PublicProfileImage:  publicProviderImage,
		PrivateProfileImage: publicProviderImage,
	}

	user, err := a.Database.CreateUser(a.ctx, arg)

	if err != nil {
		return nil, UserErrorHandler(err)
	}

	// get user object

	rsp, e := a.CreateUserObjectForAuth(user, fcmToken)

	if e != nil {
		return nil, e
	}

	return rsp, nil

}

func (a *AuthUtils) CreateUserWithProvider(email string, provider string, fcmToken string, image string) (*UserResponseTypeWithToken, *errorhandler.ErrorHandlerApi) {

	username := utils.RandomUserName(email)
	publicProviderImage := utils.RandomUserProfileImage()

	privateProfileImage := publicProviderImage

	if len(image) != 0 {
		privateProfileImage = image
	}

	arg := db.CreateUserParams{
		Username:            username,
		PasswordHash:        sql.NullString{String: "", Valid: false},
		Email:               email,
		Provider:            provider,
		PublicProfileImage:  publicProviderImage,
		PrivateProfileImage: privateProfileImage,
	}

	user, err := a.Database.CreateUser(a.ctx, arg)

	if err != nil {
		return nil, UserErrorHandler(err)
	}

	// get user object

	rsp, e := a.CreateUserObjectForAuth(user, fcmToken)

	if e != nil {
		return nil, e
	}

	return rsp, nil

}

func (a *AuthUtils) CreateUserAnonymousUser(fcmToken string) (*UserResponseTypeWithToken, *errorhandler.ErrorHandlerApi) {
	randomUser, err := utils.GenerateRandomUser()
	if (err) != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	arg := db.CreateUserParams{
		Username:            randomUser.Username,
		Email:               randomUser.Email,
		PasswordHash:        randomUser.PasswordHash,
		PublicProfileImage:  randomUser.PublicProfileImage,
		PrivateProfileImage: randomUser.PrivateProfileImage,
		Provider:            randomUser.Provider,
	}
	fmt.Println(arg)
	user, err := a.Database.CreateUser(a.ctx, arg)

	if (err) != nil {
		return nil, UserErrorHandler(err)
	}
	// get user object

	rsp, e := a.CreateUserObjectForAuth(user, fcmToken)

	if e != nil {
		return nil, e
	}

	return rsp, nil

}
