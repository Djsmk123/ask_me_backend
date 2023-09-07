package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	errorhandler "github.com/djsmk123/askmeapi/api/error_handler"
	responsehandler "github.com/djsmk123/askmeapi/api/response_handler"
	db "github.com/djsmk123/askmeapi/db/sqlc"

	"github.com/djsmk123/askmeapi/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "autherization" // Corrected the header name
	authorizationTypeBearer = "bearer"        // Corrected the token type
	authorizationPayloadKey = "autherization_payload"
)

// authMiddleware is a Gin middleware that performs token authentication.
func (server *Server) AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, errorhandler.ErrMissingAuthHeader)
			return
		}
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 2 || strings.ToLower(fields[0]) != authorizationTypeBearer {
			responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, errorhandler.ErrInvalidAuthHeader)
			return
		}
		accessToken := fields[1]
		errExpiredToken := token.ErrExpiredToken
		errInvalidToken := token.ErrInvalidToken

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			switch err {
			case errExpiredToken, errInvalidToken:
				responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, err)
				return
			}
			responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, errorhandler.ErrVerifiyingAuthHeader)
			return
		}

		//check token is valid by comparing from token database

		arg := db.GetJwtTokenUserIdParams{
			UserID:   int32(payload.ID),
			JwtToken: accessToken,
		}

		token, err := server.database.GetJwtTokenUserId(ctx, arg)
		fmt.Println(err)

		if err != nil {
			if err == sql.ErrNoRows {
				responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, errInvalidToken)
				return
			}
			responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, err)
			return
		}
		if token.ExpiresAt.Before(time.Now()) {
			responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, errExpiredToken)
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
