package utils

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	db "github.com/djsmk123/askmeapi/db/sqlc"
)

func GenerateRandomUser() (*db.CreateUserParams, error) {
	var params db.CreateUserParams
	email := RandomEmail()
	username := RandomUserName(email)
	params = db.CreateUserParams{
		Username:            username,
		Email:               email,
		PasswordHash:        sql.NullString{String: "", Valid: false},
		PublicProfileImage:  RandomUserProfileImage(),
		PrivateProfileImage: RandomUserProfileImage(),
		Provider:            "Anonymous",
	}

	return &params, nil
}

var domains = []string{"askmehelp.com", "askme1.com", "asmke.in", "askme.live", "tempaskme.net"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomEmail() string {
	randomDomain := domains[rand.Intn(len(domains))]

	// Generate a random username with a random length between 6 and 12 characters
	usernameLength := rand.Intn(7) + 6
	username := make([]byte, usernameLength)
	for i := 0; i < usernameLength; i++ {
		username[i] = byte(rand.Intn(26) + 97) // ASCII lowercase letters
	}

	fakeEmail := fmt.Sprintf("%s@%s", string(username), randomDomain)
	return fakeEmail
}

func CheckIsAnonymousUser(email string) bool {
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

func RandomUserName(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func RandomUserProfileImage() string {
	imageIndex := rand.Intn(49) + 1
	url := fmt.Sprintf("https://xsgames.co/randomusers/assets/avatars/pixel/%d.jpg", imageIndex)
	return url
}
