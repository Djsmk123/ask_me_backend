package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/djsmk123/askmeapi/token"
	"github.com/gin-gonic/gin"
)

const (
	autherizationHeaderKey  = "autherization"
	autherizationTypeBearer = "bearer"
	autherizationPayloadKey = "autherization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(autherizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("missing authorization header")
			ResponseHandlerJson(ctx, http.StatusUnauthorized, err, nil)
			return
		}
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 2 || strings.ToLower(fields[0]) != autherizationTypeBearer {
			err := errors.New("invalid authorization header")
			ResponseHandlerJson(ctx, http.StatusUnauthorized, err, nil)

			return
		}
		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			fmt.Printf("Error verifying token")
			ResponseHandlerJson(ctx, http.StatusUnauthorized, err, nil)
			return
		}
		ctx.Set(autherizationPayloadKey, payload)
		ctx.Next()

	}
}
