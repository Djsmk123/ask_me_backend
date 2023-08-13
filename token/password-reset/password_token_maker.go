package passwordreset

import (
	"time"
)

type PasswordPayloadMaker interface {
	CreateToken(id int64, duration time.Duration) (string, error)
	VerifyToken(token string) (*PasswordPayloads, error)
}
