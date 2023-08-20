package api

import (
	"net/http"
	"strings"

	responsehandler "github.com/djsmk123/askmeapi/api/response_handler"

	"github.com/djsmk123/askmeapi/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "autherization" // Corrected the header name
	authorizationTypeBearer = "bearer"        // Corrected the token type
	authorizationPayloadKey = "autherization_payload"
)

// authMiddleware is a Gin middleware that performs token authentication.
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, errMissingAuthHeader)
			return
		}
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 2 || strings.ToLower(fields[0]) != authorizationTypeBearer {
			responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, errInvalidAuthHeader)
			return
		}
		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			switch err {
			case token.ErrInvalidToken, token.ErrExpiredToken:
				responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, err)
				return
			}
			responsehandler.ResponseHandlerAbort(ctx, http.StatusUnauthorized, errVerifiyingAuthHeader)
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
