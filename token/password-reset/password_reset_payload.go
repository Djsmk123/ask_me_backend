package passwordreset

import (
	"time"

	"github.com/djsmk123/askmeapi/token"
)

type PasswordPayloads struct {
	Id        int64     `json:"id,omitempty"`
	IssuedAt  time.Time `json:"issued"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPasswordPayload(id int64, duration time.Duration) *PasswordPayloads {
	return &PasswordPayloads{
		Id:        id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (payload *PasswordPayloads) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return token.ErrExpiredToken
	}
	return nil
}
