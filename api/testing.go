package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) testing(ctx *gin.Context) {
	ResponseHandlerJson(ctx, http.StatusOK, nil, gin.H{
		"testing": true,
	})
}
