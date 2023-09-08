package auth

import (
	"context"

	db "github.com/djsmk123/askmeapi/db/sqlc"
	"github.com/djsmk123/askmeapi/token"
	passwordreset "github.com/djsmk123/askmeapi/token/password-reset"
	"github.com/djsmk123/askmeapi/utils"
)

type AuthUtils struct {
	Config          utils.Config
	Database        db.DBExec
	TokenMaker      token.Maker
	ctx             context.Context
	PasswordRequest passwordreset.PasswordPayloadMaker
}

func NewAuthUtils(config utils.Config, db db.DBExec, tokenMaker token.Maker, ctx context.Context, pass passwordreset.PasswordPayloadMaker) AuthUtils {
	return AuthUtils{
		Config:          config,
		Database:        db,
		TokenMaker:      tokenMaker,
		ctx:             ctx,
		PasswordRequest: pass,
	}
}
