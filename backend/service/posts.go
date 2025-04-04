package service

import (
	"net/http"

	sql "goblogeasyg/sql/model"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var artical map[string]interface{}
	if err := c.ShouldBind(&artical); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if artical["content"] == "" || artical["title"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title or content cannot be empty"})
		return
	}

	// 获取tag并构造sql.Tag类型结构体
	tagsData, ok := artical["tags"].([]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tags format"})
		return
	}
	var tags []sql.Tag
	for _, t := range tagsData {
		tags = append(tags, sql.Tag{Name: t.(string)})
	}

	err := sql.CreatePost(sql.Article{
		Content: artical["content"].(string),
		Title:   artical["title"].(string),
		Tags:    tags,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "create"})
}

func GetPosts(c *gin.Context) {
	posts, err := sql.GetPosts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"posts": posts})
}

func DBinit(c *gin.Context) {
	sql.AutoMigrate()
	c.JSON(200, gin.H{"message": "init db success"})
}
