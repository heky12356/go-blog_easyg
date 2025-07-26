package middleware

import (
	"net/http"
	"os"
	"strings"

	"goblogeasyg/internal/cache"
	"goblogeasyg/internal/response"
	utils "goblogeasyg/internal/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = os.Getenv("JWT_KEY")

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		if tokenString == "" {
			response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// 检查token是否在黑名单中
		if cache.IsInBlacklist(tokenString) {
			response.ErrorResponse(c, http.StatusUnauthorized, "token已失效")
			c.Abort()
			return
		}

		var claims utils.Claims
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNotSupported
			}
			return []byte(jwtKey), nil
		})
		if err != nil {
			response.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if !token.Valid {
			response.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if claims.Type != "access" {
			response.ErrorResponse(c, http.StatusUnauthorized, "Invalid token type")
			c.Abort()
			return
		}

		c.Next()
	}
}
