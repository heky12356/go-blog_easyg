package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = os.Getenv("JWT_KEY")

type Claims struct {
	Username string `json:"username"`
	Type     string `json:"token_type"`
	jwt.StandardClaims
}

func CreateRefreshToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
		Type:     "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 设置过期时间
		},
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateAssessToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
		Type:     "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // 设置过期时间
		},
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
