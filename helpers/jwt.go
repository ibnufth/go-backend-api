package helpers

import (
	"ibnufth/backend-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) string {
	// Set expiration time for the token, 360 minutes from now
	expirationTime := time.Now().Add(360 * time.Minute)

	// create the JWT claims, which includes the username and expiry time
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(config.GetJWTSecret())

	return token
}
