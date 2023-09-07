package auth

import (
	db "github.com/djsmk123/askmeapi/db/sqlc"
)

func (a *AuthUtils) CreateFcmToken(token string, user db.User, jwtId int, database db.DBExec) error {
	if token == "Null" || len(token) <= 0 {
		return nil
	}
	//save token in the database

	arg := db.CreateFcmTokenParams{
		ID:       int32(jwtId),
		UserID:   user.ID,
		FcmToken: token,
		IsValid:  true,
	}
	_, err := a.Database.CreateFcmToken(a.ctx, arg)

	if err != nil {
		return err
	}
	return nil
}
