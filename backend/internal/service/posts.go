package service

import (
	"net/http"

	sql "goblogeasyg/internal/sql/model"
	"goblogeasyg/internal/utils"

	"github.com/gin-gonic/gin"
)

// 创建
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

	// 创建uid
	uid, err := utils.CreateUid()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 使用sql.CreatePost插入数据
	err = sql.CreatePost(sql.Article{
		Content: artical["content"].(string),
		Title:   artical["title"].(string),
		Tags:    tags,
		Uid:     uid,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "create"})
}

// 获取
func GetPosts(c *gin.Context) {
	posts, err := sql.GetPostsBase()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	uid := c.Param("uid")
	post, err := sql.GetPostByUid(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

// 删除
func DeletePost(c *gin.Context) {
	uid := c.Param("uid")
	err := sql.DeletePost(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "delete success"})
}
