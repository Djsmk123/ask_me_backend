package responsehandler

import (
	"github.com/gin-gonic/gin"
)

type APIRESPONSE struct {
	StatusCode int64       `json:"status_code"`
	Message    string      `json:"message"`
	Data       SuccessData `json:"data"`
	Status     bool        `json:"status"`
}

type SuccessData interface{}

func ResponseHandlerAbort(ctx *gin.Context, code int64, err error) {
	response := APIRESPONSE{
		StatusCode: code,
		Message:    err.Error(),
		Data:       nil,
		Status:     false,
	}
	ctx.AbortWithStatusJSON(int(code), response)

}
func ResponseHandlerJson(ctx *gin.Context, code int64, err error, data SuccessData) {
	var response APIRESPONSE

	if err != nil {
		response = APIRESPONSE{
			StatusCode: code,
			Message:    err.Error(),
			Data:       nil,
			Status:     false,
		}
		ctx.JSON(int(code), response)
		return
	}
	response = APIRESPONSE{
		StatusCode: code,
		Message:    "Success",
		Data:       data,
		Status:     true,
	}
	ctx.JSON(int(code), response)
}
