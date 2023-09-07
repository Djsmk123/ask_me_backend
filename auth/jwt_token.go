package auth

import (
	"time"

	db "github.com/djsmk123/askmeapi/db/sqlc"
)

func (a *AuthUtils) CreateJwtToken(userId int64, username string) (string, error) {
	accesstoken, err := a.TokenMaker.CreateToken(userId, username, a.Config.AccessTokenDuration)

	if err != nil {
		return "", err
	}
	payload, err := a.TokenMaker.VerifyToken(accesstoken)
	if err != nil {
		return "", err
	}
	//save token to database

	arg := db.CreateJwtTokenParams{
		UserID:    int32(userId),
		JwtToken:  accesstoken,
		CreatedAt: payload.IssuedAt,
		ExpiresAt: payload.ExpiredAt,
	}

	__, err := a.Database.CreateJwtToken(a.ctx, arg)

	if err != nil || __.ExpiresAt.Before(time.Now()) {
		return "", err
	}
	return accesstoken, nil

}
