package api

import (
	"context"
	"fmt"

	"github.com/djsmk123/askmeapi/auth"
	db "github.com/djsmk123/askmeapi/db/sqlc"
	passwordreset "github.com/djsmk123/askmeapi/token/password-reset"
	"github.com/djsmk123/askmeapi/utils"
	"github.com/gin-gonic/gin"

	"github.com/djsmk123/askmeapi/token"
)

type Server struct {
	Config        utils.Config
	database      db.DBExec
	Router        *gin.Engine
	TokenMaker    token.Maker
	PasswordReset passwordreset.PasswordPayloadMaker
	Auth          auth.AuthUtils
}

func NewServer(config utils.Config, store db.DBExec) (*Server, error) {
	tokenMaker, err := token.NewJwtMaker(config.TokkenStructureKey)
	passwordResetMaker := passwordreset.NewPassWordResetMaker(config.TokkenStructureKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		Config:        config,
		database:      store,
		TokenMaker:    tokenMaker,
		PasswordReset: passwordResetMaker,
		Auth:          auth.NewAuthUtils(config, store, tokenMaker, context.Background()),
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {

	return server.Router.Run(address)
}
