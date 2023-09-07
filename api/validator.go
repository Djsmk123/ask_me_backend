package api

import (
	"net/http"

	responsehandler "github.com/djsmk123/askmeapi/api/response_handler"
	"github.com/gin-gonic/gin"
)

type Response struct{}

func JsonBinder(ctx *gin.Context, req interface{}) interface{} {
	var res *Response
	if err := ctx.ShouldBindJSON(&res); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return nil
	}
	return res

}
