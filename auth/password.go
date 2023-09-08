package auth

import (
	"database/sql"
	"net/http"
	"net/url"
	"strings"
	"time"

	errorhandler "github.com/djsmk123/askmeapi/api/error_handler"
	db "github.com/djsmk123/askmeapi/db/sqlc"
	"github.com/djsmk123/askmeapi/utils"
)

type PasswordRequestResponse struct {
	Message string          `json:"message"`
	Data    utils.EmailData `json:"data"`
}
type PasswordResetResponse struct {
	Message string `json:"message"`
}

func (a *AuthUtils) RequestResetPassword(email string) (*PasswordRequestResponse, *errorhandler.ErrorHandlerApi) {
	// Check if user is not anonymous.
	if isAnonymous := utils.CheckIsAnonymousUser(email); isAnonymous {
		return nil, &errorhandler.ErrorHandlerApi{
			Error:      ErrAnonymousUserFound,
			StatusCode: http.StatusBadRequest,
		}
	}

	// Find user by email to check if the user exists or not.
	user, err := a.Database.GetUserByEmail(a.ctx, email)
	if err != nil {
		return nil, UserErrorHandler(err)
	}

	// Check if the user provider is email-password or not.
	if strings.ToLower(user.Provider) != "password" && !user.PasswordHash.Valid {
		return nil, &errorhandler.ErrorHandlerApi{
			Error:      ErrAnonymousUserFound,
			StatusCode: http.StatusBadRequest,
		}
	}

	resetToken, err := a.PasswordRequest.CreateToken(int64(user.ID), 10*time.Minute)
	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}

	encryptedUrl := url.QueryEscape(resetToken)
	emailData := utils.EmailData{
		URL:       a.Config.BaseUrl + "reset-password-page?token=" + encryptedUrl,
		FirstName: user.Email,
		Subject:   "Your password reset token (valid for 10min)",
	}

	err = utils.SendEmail(user.Email, &a.Config, &emailData, "resetPassword.html")
	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}

	return &PasswordRequestResponse{
		Message: "Email has been sent successfully",
		Data:    emailData,
	}, nil
}
func (a *AuthUtils) ResetPassword(password string, token string) (*PasswordResetResponse, *errorhandler.ErrorHandlerApi) {
	payload, err := a.PasswordRequest.VerifyToken(token)
	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	err = payload.Valid()
	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	newPasswordHash, err := utils.HashPassword(password)
	if err != nil {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}

	user, err := a.Database.UpdateUserPassword(a.ctx, db.UpdateUserPasswordParams{
		ID: int32(payload.Id),
		PasswordHash: sql.NullString{
			String: newPasswordHash,
			Valid:  true,
		},
	})
	if err != nil || user.ID < 0 {
		return nil, errorhandler.InternalServerErrorHandler(err)
	}
	return &PasswordResetResponse{
		Message: "Password reset successfully",
	}, nil

}
