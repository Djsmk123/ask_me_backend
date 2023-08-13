package passwordreset

import (
	"fmt"
	"time"
)

type PasswordResetToken struct {
	Token string `json:"token"`
}

func NewPassWordResetMaker(secret string) *PasswordResetToken {
	return &PasswordResetToken{
		Token: secret,
	}
}

func (maker *PasswordResetToken) CreateToken(id int64, duration time.Duration) (string, error) {
	key := []byte(maker.Token)
	payload := NewPasswordPayload(id, duration)
	encryptedData, err := encryptStruct(*payload, key)
	if err != nil {
		fmt.Println("Error encrypting struct:", err)
		return "", err

	}
	return encryptedData, nil
}
func (maker *PasswordResetToken) VerifyToken(token string) (*PasswordPayloads, error) {
	key := []byte(maker.Token)

	decryptedData, err := decryptStruct(token, key)
	if err != nil {
		fmt.Println("Error decrypting struct:", err)
		return nil, err
	}
	return &decryptedData, nil

}
