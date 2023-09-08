package responsehandler

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	StatusCode int64       `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Success    bool        `json:"success"`
}

func ResponseHandlerAbort(ctx *gin.Context, code int64, err error) {
	response := APIResponse{
		StatusCode: code,
		Message:    err.Error(),
		Data:       nil,
		Success:    false,
	}
	ctx.AbortWithStatusJSON(int(code), response)
}

func ResponseHandlerJSON(ctx *gin.Context, code int64, err error, data interface{}) {
	response := APIResponse{
		StatusCode: code,
		Data:       data,
	}

	if err != nil {
		response.Success = false
		message := err.Error()

		if data != nil {
			response.Data = data
		}

		if message == "dial tcp [::1]:5432: connectex: No connection could be made because the target machine actively refused it." {
			message = "Connection establishment failed"
		}

		response.Message = message
	} else {
		response.Success = true
		response.Message = "Request has been served successfully"
	}

	ctx.JSON(int(code), response)
}
