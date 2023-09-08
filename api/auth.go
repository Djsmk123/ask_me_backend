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

// CreateNewUserRequest represents the request body for creating a new user.
type CreateNewUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FcmToken string `json:"fcm_token"`
}

// CreateAnoUserRequest represents the request body for creating an anonymous user.
type CreateAnoUserRequest struct {
	FcmToken string `json:"fcm_token"`
}

// CreateAnonymousUser creates an anonymous user.
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

// CreateUser creates a new user.
func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateNewUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJSON(ctx, http.StatusBadRequest, err, nil)
		return
	}

	rsp, err := server.Auth.CreateUser(req.Email, req.Password, req.FcmToken)

	RequestHandlerJSON(ctx, rsp, err)
}

// LoginUserRequest represents the request body for user login.
type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FcmToken string `json:"fcm_token"`
}

// LoginUser handles user login.
func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJSON(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	rsp, err := server.Auth.Login(req.Email, req.Password, req.FcmToken)

	RequestHandlerJSON(ctx, rsp, err)
}

// RequestHandlerJSON handles JSON responses.
func RequestHandlerJSON(ctx *gin.Context, rsp interface{}, err *errorhandler.ErrorHandlerApi) {
	if err != nil {
		responsehandler.ResponseHandlerJSON(ctx, int64(err.StatusCode), err.Error, nil)
		return
	}
	responsehandler.ResponseHandlerJSON(ctx, http.StatusOK, nil, rsp)
}

// SocialLoginRequestType represents the request body for social login.
type SocialLoginRequestType struct {
	Email               string `json:"email" binding:"required,email"`
	PrivateProfileImage string `json:"private_profile_image"`
	Provider            string `json:"provider" binding:"required"`
	FcmToken            string `json:"fcm_token"`
}

// SocialLogin handles social login.
func (server *Server) SocialLogin(ctx *gin.Context) {
	var req SocialLoginRequestType
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJSON(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	rsp, err := server.Auth.SocialLogin(req.Email, req.Provider, req.FcmToken, req.PrivateProfileImage)

	RequestHandlerJSON(ctx, rsp, err)
}

// GetUser gets user information.
func (server *Server) GetUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	rsp, err := server.Auth.GetUser(int32(authPayload.ID))
	RequestHandlerJSON(ctx, rsp, err)
}

// DeleteUser deletes a user.
func (server *Server) DeleteUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	rsp, err := server.Auth.DeleteUser(int32(authPayload.ID))
	RequestHandlerJSON(ctx, rsp, err)
}

// PasswordResetRequest represents the request body for password reset request.
type PasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// PasswordResetRequest sends a password reset request.
func (server *Server) PasswordResetRequest(ctx *gin.Context) {
	var req PasswordResetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJSON(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	rsp, err := server.Auth.RequestResetPassword(req.Email)
	RequestHandlerJSON(ctx, rsp, err)
}

// ResetPasswordInput represents the request body for resetting a password.
type ResetPasswordInput struct {
	Password string `json:"password" binding:"required"`
}

// ResetPaswordVerify verifies and resets the password.
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

// LogoutUser logs out the user.
func (server *Server) LogoutUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	fields := strings.Fields(ctx.GetHeader("authorization"))

	payload := fields[1]

	rsp, err := server.Auth.Logout(int32(authPayload.ID), payload)
	RequestHandlerJSON(ctx, rsp, err)
}
