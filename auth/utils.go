package auth

import (
	"context"

	db "github.com/djsmk123/askmeapi/db/sqlc"
	"github.com/djsmk123/askmeapi/token"
	"github.com/djsmk123/askmeapi/utils"
)

type AuthUtils struct {
	Config     utils.Config
	Database   db.DBExec
	TokenMaker token.Maker
	ctx        context.Context
}

func NewAuthUtils(config utils.Config, db db.DBExec, tokenMaker token.Maker, ctx context.Context) AuthUtils {
	return AuthUtils{
		Config:     config,
		Database:   db,
		TokenMaker: tokenMaker,
		ctx:        ctx,
	}
}
