package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func GenerateJWT( id string) (string, error) {
	if len(secretKey) == 0 {
		log.Fatal("SECRET_KEY environment variable not set")
	}

	claims := jwt.MapClaims{

		"id":  id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
