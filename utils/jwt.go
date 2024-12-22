package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/rizqullorayhan/go-fiber-gorm/config"
)

var SecretKey = []byte(config.Config("JWT_SECRET_KEY"))

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString(SecretKey)
}