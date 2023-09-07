package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	responsehandler "github.com/djsmk123/askmeapi/api/response_handler"
	"github.com/djsmk123/askmeapi/auth"
	db "github.com/djsmk123/askmeapi/db/sqlc"
	"github.com/djsmk123/askmeapi/token"
	"github.com/djsmk123/askmeapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateNewUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FcmToken string `json:"fcm_token"`
}

type CreateAnoUserRequest struct {
	FcmToken string `json:"fcm_token"`
}

func (server *Server) CreateAnonymousUser(ctx *gin.Context) {
	var req CreateAnoUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	rsp, err := server.Auth.CreateUserAnonymousUser(req.FcmToken)

	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, int64(err.StatusCode), err.Error, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, rsp)

}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateNewUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusBadRequest, err, nil)
		return
	}

	rsp, err := server.Auth.CreateUser(req.Email, req.Password, req.FcmToken)

	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, int64(err.StatusCode), err.Error, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, rsp)

}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FcmToken string `json:"fcm_token"`
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	rsp, err := server.Auth.Login(req.Email, req.Password, req.FcmToken)

	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, int64(err.StatusCode), err.Error, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, rsp)

}

type SocialLoginRequestType struct {
	Email               string `json:"email" binding:"required,email"`
	PrivateProfileImage string `json:"private_profile_image"`
	Provider            string `json:"provider" binding:"required"`
	FcmToken            string `json:"fcm_token"`
}

func (server *Server) SocialLogin(ctx *gin.Context) {
	var req SocialLoginRequestType
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	user, err := server.database.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows || err.Error() == "no rows in result set" {
			profileImage := req.PrivateProfileImage
			if (len(profileImage)) == 0 {
				profileImage = utils.RandomUserProfileImage()
			}
			arg := db.CreateUserParams{
				Username:            utils.RandomUserName(req.Email),
				Provider:            req.Provider,
				Email:               req.Email,
				PasswordHash:        sql.NullString{String: "", Valid: false},
				PublicProfileImage:  utils.RandomUserProfileImage(),
				PrivateProfileImage: profileImage,
			}
			user, err := server.database.CreateUser(ctx, arg)

			if err != nil {
				responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
				return
			}
			server.Auth.CreateUserObjectForAuth(user, req.FcmToken)

			return

		}

		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	//check if user is not anonymous or password based

	provider := strings.ToLower(user.Provider)

	if provider != "anonymous" && provider != "password" {
		responsehandler.ResponseHandlerJson(ctx, http.StatusBadRequest, errors.New("invalid resource request"), nil)
		return
	}
	server.Auth.CreateUserObjectForAuth(user, req.FcmToken)
}

func (server *Server) GetUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.database.GetUserByID(ctx, int32(authPayload.ID))

	if (err) != nil {
		UserErrorHandler(ctx, err)
		return
	}
	server.Auth.CreateUserObjectWithoutToken(user)
}
func (server *Server) DeleteUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	//delete question where user_id =id
	server.database.DeleteAnswerByUserId(ctx, int32(authPayload.ID))

	server.database.DeleteQuestionByUserId(ctx, int32(authPayload.ID))

	server.database.DeleteFcmTokenByUserId(ctx, int32(authPayload.ID))
	server.database.DeleteJWTokenByUserId(ctx, int32(authPayload.ID))

	_, err := server.database.DeleteUserById(ctx, int32(authPayload.ID))
	if (err) != nil {
		UserErrorHandler(ctx, err)
		return
	}

	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, "delete-successfully")

}

func UserErrorHandler(ctx *gin.Context, err error) {
	if err == sql.ErrNoRows {
		responsehandler.ResponseHandlerJson(ctx, http.StatusNotFound, auth.ErrUserNotExist, nil)
		return
	}
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation", "unique_key_voilation":
			responsehandler.ResponseHandlerJson(ctx, http.StatusForbidden, auth.ErrUserAlreadyExist, nil)
			return
		}
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
}

type PasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (server *Server) PasswordResetRequest(ctx *gin.Context) {
	var req PasswordResetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	// check if user is not anonymous
	isAnonymous := utils.CheckIsAnonymousUser(req.Email)

	if isAnonymous {
		responsehandler.ResponseHandlerJson(ctx, http.StatusBadRequest, auth.ErrAnonymousUserFound, nil)
		return
	}

	//find user by email to check if user exists or not

	user, err := server.database.GetUserByEmail(ctx, req.Email)

	if (err) != nil {
		UserErrorHandler(ctx, err)
		return
	}
	// check if user is not logged with social account
	if !user.PasswordHash.Valid {
		responsehandler.ResponseHandlerJson(ctx, http.StatusBadRequest, errors.New("invalid resource request"), nil)
	}

	resetToken, err := server.PasswordReset.CreateToken(int64(user.ID), 10*time.Minute)
	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	encryptedUrl := url.QueryEscape(resetToken)
	emailData := utils.EmailData{
		URL:       server.Config.BaseUrl + "reset-password-page?token=" + encryptedUrl,
		FirstName: user.Email,
		Subject:   "Your password reset token (valid for 10min)",
	}

	err = utils.SendEmail(user.Email, &server.Config, &emailData, "resetPassword.html")
	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, gin.H{
		"message":  "Email has been sent successfully",
		"response": emailData,
	})

}

type ResetPasswordInput struct {
	Password string `json:"password" binding:"required"`
}

func (server *Server) ResetPaswordVerify(ctx *gin.Context) {
	var req ResetPasswordInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	token, exists := ctx.GetQuery("token")

	if !exists && token != "" {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, errors.New("invalid url"), nil)
		return
	}

	payload, err := server.PasswordReset.VerifyToken(token)
	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	err = payload.Valid()
	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusUnauthorized, err, nil)
		return
	}
	newPasswordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	user, err := server.database.UpdateUserPassword(ctx, db.UpdateUserPasswordParams{
		ID: int32(payload.Id),
		PasswordHash: sql.NullString{
			String: newPasswordHash,
			Valid:  true,
		},
	})
	if err != nil || user.ID < 0 {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, gin.H{"status": "success", "message": "Password data updated successfully"})
}

func (server *Server) LogoutUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	fields := strings.Fields(ctx.GetHeader("authorization"))

	payload := fields[1]

	arg := db.UpdateJwtTokenParams{
		JwtToken: payload,
		UserID:   int32(authPayload.ID),
	}
	_, err := server.database.UpdateJwtToken(ctx, arg)

	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	//update fcm toke if exist
	arg1 := db.UpdateFcmTokenParams{
		ID:      int32(authPayload.ID),
		IsValid: false,
	}
	_, err = server.database.UpdateFcmToken(ctx, arg1)

	if err != nil && err != sql.ErrNoRows {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, "Successfully log-out")

}
