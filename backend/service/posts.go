package service

import "github.com/gin-gonic/gin"

func CreatePost(c *gin.Context) {
	c.JSON(200, gin.H{"message": "create"})
}
