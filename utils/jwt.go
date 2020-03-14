package utils

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT utils
func GenerateJWT(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte("argadev123"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
