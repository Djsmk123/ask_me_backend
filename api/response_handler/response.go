package responsehandler

import (
	"github.com/gin-gonic/gin"
)

type APIRESPONSE struct {
	StatusCode int64       `json:"status_code"`
	Message    string      `json:"message"`
	Data       SuccessData `json:"data"`
	Success    bool        `json:"success"`
}

type SuccessData interface{}

func ResponseHandlerAbort(ctx *gin.Context, code int64, err error) {
	response := APIRESPONSE{
		StatusCode: code,
		Message:    err.Error(),
		Data:       nil,
		Success:    false,
	}
	ctx.AbortWithStatusJSON(int(code), response)

}
func ResponseHandlerJson(ctx *gin.Context, code int64, err error, data SuccessData) {
	var response APIRESPONSE

	if err != nil {
		var response SuccessData = "invalid request"
		e := err.Error()
		if data != nil {
			response = data
		}
		if e == "dial tcp [::1]:5432: connectex: No connection could be made because the target machine actively refused it." {
			e = "Connection establishment failed"
		}
		response = APIRESPONSE{
			StatusCode: code,
			Message:    e,
			Data:       response,
			Success:    false,
		}
		ctx.JSON(int(code), response)
		return
	}
	response = APIRESPONSE{
		StatusCode: code,
		Message:    "Request has been served successfully",
		Data:       data,
		Success:    true,
	}
	ctx.JSON(int(code), response)
}
