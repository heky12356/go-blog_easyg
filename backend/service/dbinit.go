package service

import (
	"net/http"

	sql "goblogeasyg/sql/model"

	"github.com/gin-gonic/gin"
)

// 初始化
func DBinit(c *gin.Context) {
	err := sql.AutoMigrateArticle()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = sql.AutoMigrateUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "init db success"})
}
