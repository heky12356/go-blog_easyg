package utils

import "github.com/golang-jwt/jwt"

func RefreshToken(token string) (string, error) {
	var claims Claims
	refreshtoken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return "", err
	}
	if !refreshtoken.Valid {
		return "", err
	}
	if claims.Type != "refresh" {
		return "", err
	}
	accessToken, err := CreateRefreshToken(claims.Username)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
