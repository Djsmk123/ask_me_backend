package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	db "github.com/djsmk123/askmeapi/db/sqlc"
)

func GenerateRandomUser() (*db.CreateUserParams, error) {
	var params db.CreateUserParams
	email := RandomEmail()
	username, err := RandomUserName(email)
	if err != nil {
		return &params, err
	}
	params = db.CreateUserParams{
		Username:     username,
		Email:        email,
		PasswordHash: sql.NullString{String: "", Valid: false},
	}

	return &params, nil
}

var domains = []string{"askmehelp.com", "askme1.com", "asmke.in", "askme.live", "tempaskme.net"}

func RandomEmail() string {

	randomDomain := domains[rand.Intn(len(domains))]
	usernameLength := rand.Intn(7) + 6
	username := make([]byte, usernameLength)
	for i := 0; i < usernameLength; i++ {
		username[i] = byte(rand.Intn(26) + 97) // ASCII lowercase letters
	}

	fakeEmail := fmt.Sprintf("%s@%s", string(username), randomDomain)
	return fakeEmail
}

func checkIsAnonymousUser(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false // Invalid email format
	}

	domain := parts[1]
	return isAnonymousDomain(domain)
}
func isAnonymousDomain(domain string) bool {
	for _, anonDomain := range domains {
		if domain == anonDomain {
			return true
		}
	}
	return false
}

func RandomUserName(email string) (string, error) {
	parts := strings.Split(email, "@")
	if len(parts) > 0 {
		return parts[0], nil
	}
	return "", errors.New("invalid user name")
}
