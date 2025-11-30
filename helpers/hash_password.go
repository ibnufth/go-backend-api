package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {

	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
