package utils

import (
	"strings"

	"github.com/golang-jwt/jwt"
)

func ParseToken(tokenString string) (*Claims, error) {
	var claims Claims
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}
	if claims.Type != "access" {
		return nil, jwt.ErrInvalidKey
	}
	return &claims, nil
}
