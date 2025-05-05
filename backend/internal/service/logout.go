package service

import (
	"net/http"
	"strings"
	"time"

	"goblogeasyg/internal/cache"
	jwt "goblogeasyg/internal/utils/jwt"

	"github.com/gin-gonic/gin"
)

// Logout handles user logout by invalidating the current token
func Logout(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
		return
	}

	// 从Bearer token中提取实际的token
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// 解析token以获取过期时间
	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的token"})
		return
	}

	// 计算token剩余有效期
	expiration := time.Until(time.Unix(claims.ExpiresAt, 0))
	if expiration <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token已过期"})
		return
	}

	// 将token加入黑名单，使用剩余有效期作为过期时间
	err = cache.AddToBlacklist(tokenString, expiration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加token到黑名单失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登出成功"})
}
