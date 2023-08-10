package api

import (
	"fmt"

	db "github.com/djsmk123/askmeapi/db/sqlc"
	"github.com/djsmk123/askmeapi/utils"
	"github.com/gin-gonic/gin"

	"github.com/djsmk123/askmeapi/token"
)

type Server struct {
	config     utils.Config
	store      db.DBExec
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config utils.Config, store db.DBExec) (*Server, error) {
	tokenMaker, err := token.NewJwtMaker(config.TokkenStructureKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}
func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/", server.index)
	v1 := router.Group("/api/v1")

	authRoutesV1 := v1.Use(authMiddleware(server.tokenMaker))
	v1.POST("/create-user", server.CreateUser)
	v1.POST("/create-ano-user", server.CreateAnonymousUser)
	v1.POST("/login-user", server.LoginUser)
	v1.POST("/social-login", server.SocialLogin)

	authRoutesV1.GET("/delete-user/", server.DeleteUser)

	authRoutesV1.POST("/create-question", server.CreateQuestion)
	authRoutesV1.POST("/update-question", server.UpdateQuestionById)
	authRoutesV1.GET("/delete-question/:id", server.DeleteQuestionById)
	authRoutesV1.GET("/questions", server.ListQuestion)
	authRoutesV1.GET("/question/:id", server.GetQuestionByID)

	authRoutesV1.POST("/create-answer", server.CreateAnswer)
	authRoutesV1.POST("/update-answer", server.UpdateAnswerById)
	authRoutesV1.GET("/delete-answer/:id", server.DeleteAnswerById)
	authRoutesV1.GET("/answers", server.ListAnswers)
	authRoutesV1.GET("/answer/:id", server.GetAnswerByID)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

type APIRESPONSE struct {
	StatusCode int64       `json:"status_code"`
	Message    string      `json:"message"`
	Data       SuccessData `json:"data"`
	Status     bool        `json:"status"`
}

type SuccessData interface{}

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
