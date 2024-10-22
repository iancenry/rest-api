package handler

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	return tokenString, err
}

// JWT token is simply a signed JSON object
