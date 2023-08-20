package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/djsmk123/askmeapi/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "autherization" // Corrected the header name
	authorizationTypeBearer = "bearer"        // Corrected the token type
	authorizationPayloadKey = "autherization_payload"
)

var (
	errMissingAuthHeader    = errors.New("missing authorization header")
	errInvalidAuthHeader    = errors.New("invalid authorization header")
	errVerifiyingAuthHeader = errors.New("invalid authorization key")
)

// authMiddleware is a Gin middleware that performs token authentication.
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ResponseHandlerJson(ctx, http.StatusUnauthorized, errMissingAuthHeader, nil)
			return
		}
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 2 || strings.ToLower(fields[0]) != authorizationTypeBearer {
			ResponseHandlerJson(ctx, http.StatusUnauthorized, errInvalidAuthHeader, nil)
			return
		}
		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ResponseHandlerJson(ctx, http.StatusUnauthorized, errVerifiyingAuthHeader, nil)
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
