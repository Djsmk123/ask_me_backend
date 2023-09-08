package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	errorhandler "github.com/djsmk123/askmeapi/api/error_handler"
	responsehandler "github.com/djsmk123/askmeapi/api/response_handler"
	"github.com/djsmk123/askmeapi/token"
	"github.com/gin-gonic/gin"
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
		responsehandler.ResponseHandlerJSON(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	rsp, err := server.Auth.CreateUserAnonymousUser(req.FcmToken)

	RequestHandlerJSON(ctx, rsp, err)

}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateNewUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJSON(ctx, http.StatusBadRequest, err, nil)
		return
	}

	rsp, err := server.Auth.CreateUser(req.Email, req.Password, req.FcmToken)

	RequestHandlerJSON(ctx, rsp, err)

}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FcmToken string `json:"fcm_token"`
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJSON(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	rsp, err := server.Auth.Login(req.Email, req.Password, req.FcmToken)
	RequestHandlerJSON(ctx, rsp, err)
}
func RequestHandlerJSON(ctx *gin.Context, rsp interface{}, err *errorhandler.ErrorHandlerApi) {
	if err != nil {
		responsehandler.ResponseHandlerJSON(ctx, int64(err.StatusCode), err.Error, nil)
		return
	}
	responsehandler.ResponseHandlerJSON(ctx, http.StatusOK, nil, rsp)
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
		responsehandler.ResponseHandlerJSON(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	rsp, err := server.Auth.SocialLogin(req.Email, req.Provider, req.FcmToken, req.PrivateProfileImage)
	RequestHandlerJSON(ctx, rsp, err)

}

func (server *Server) GetUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	rsp, err := server.Auth.GetUser(int32(authPayload.ID))
	RequestHandlerJSON(ctx, rsp, err)
}
func (server *Server) DeleteUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	rsp, err := server.Auth.DeleteUser(int32(authPayload.ID))
	RequestHandlerJSON(ctx, rsp, err)
}

type PasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (server *Server) PasswordResetRequest(ctx *gin.Context) {
	var req PasswordResetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJSON(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	rsp, err := server.Auth.RequestResetPassword(req.Email)
	RequestHandlerJSON(ctx, rsp, err)

}

type ResetPasswordInput struct {
	Password string `json:"password" binding:"required"`
}

func (server *Server) ResetPaswordVerify(ctx *gin.Context) {
	var req ResetPasswordInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJSON(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	token, exists := ctx.GetQuery("token")

	if !exists && token != "" {
		responsehandler.ResponseHandlerJSON(ctx, http.StatusInternalServerError, errors.New("invalid url"), nil)
		return
	}
	rsp, err := server.Auth.ResetPassword(req.Password, token)
	RequestHandlerJSON(ctx, rsp, err)

}

func (server *Server) LogoutUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	fields := strings.Fields(ctx.GetHeader("authorization"))

	payload := fields[1]

	rsp, err := server.Auth.Logout(int32(authPayload.ID), payload)
	RequestHandlerJSON(ctx, rsp, err)

}
