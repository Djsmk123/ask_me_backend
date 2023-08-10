package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) index(ctx *gin.Context) {
	ResponseHandlerJson(ctx, http.StatusOK, nil, struct{}{})
}
