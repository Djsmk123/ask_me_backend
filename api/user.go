package api

import (
	"database/sql"
	"fmt"

	"net/http"
	"time"

	db "github.com/djsmk123/askmeapi/db/sqlc"
	"github.com/djsmk123/askmeapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type UserResponseType struct {
	ID        int32     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UserResponseTypeWithToken struct {
	AccessToken string           `json:"token"`
	User        UserResponseType `json:"user"`
}
type CreateNewUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func GetUserResponse(user db.User) UserResponseType {
	return UserResponseType{
		ID:        user.ID,
		Username:  user.Email,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt,
	}
}

func (server *Server) CreateAnonymousUser(ctx *gin.Context) {
	randomUser, err := utils.GenerateRandomUser()
	if (err) != nil {
		ResponseHandlerJson(ctx, http.StatusBadRequest, err, nil)
		return
	}
	arg := db.CreateUserParams{
		Username:     randomUser.Username,
		Email:        randomUser.Email,
		PasswordHash: randomUser.PasswordHash,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation", "unique_key_voilation":
				ResponseHandlerJson(ctx, http.StatusForbidden, errUserAlreadyExist, nil)
				return
			}

		}
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)

		return
	}
	accesstoken, err := server.tokenMaker.CreateToken(int64(user.ID), user.Username, server.config.AccessTokenDuration)

	if err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	resp := GetUserResponse(user)

	newResp := UserResponseTypeWithToken{
		AccessToken: accesstoken,
		User:        resp,
	}
	ResponseHandlerJson(ctx, http.StatusOK, nil, newResp)

}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateNewUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ResponseHandlerJson(ctx, http.StatusBadRequest, err, nil)
		return
	}
	hashPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	arg := db.CreateUserParams{
		Username:     req.Username,
		PasswordHash: sql.NullString{String: hashPassword, Valid: true},
		Email:        req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation", "unique_key_voilation":
				ResponseHandlerJson(ctx, http.StatusForbidden, errUserAlreadyExist, nil)
				return
			}

		}
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)

		return
	}
	accesstoken, err := server.tokenMaker.CreateToken(int64(user.ID), user.Username, server.config.AccessTokenDuration)

	if err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	resp := GetUserResponse(user)

	newResp := UserResponseTypeWithToken{
		AccessToken: accesstoken,
		User:        resp,
	}
	ResponseHandlerJson(ctx, http.StatusOK, nil, newResp)
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6" `
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("Error getting user", err == sql.ErrNoRows)
		if err == sql.ErrNoRows || err.Error() == "no rows in result set" {
			ResponseHandlerJson(ctx, http.StatusNotFound, errUserNotExist, nil)
			return
		}
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	err = utils.CheckPassword(req.Password, user.PasswordHash.String)
	if err != nil {
		ResponseHandlerJson(ctx, http.StatusUnauthorized, errWrongPassword, nil)
		return
	}
	accesstoken, err := server.tokenMaker.CreateToken(int64(user.ID), user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	rsp := UserResponseTypeWithToken{
		AccessToken: accesstoken,
		User:        GetUserResponse(user),
	}
	ResponseHandlerJson(ctx, http.StatusOK, nil, rsp)

}
