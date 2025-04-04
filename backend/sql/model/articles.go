package sql

import (
	"fmt"

	"goblogeasyg/sql"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name    string
	Aticles []Article `gorm:"many2many:article_tags;"`
}
type Article struct {
	gorm.Model
	Content string
	Title   string
	Tags    []Tag `gorm:"many2many:article_tags;"`
}

func AutoMigrate() (err error) {
	db := sql.GetDB()
	err = db.AutoMigrate(&Article{}, &Tag{})
	return
}

func CreatePost(artical Article) (err error) {
	db := sql.GetDB()
	if artical.Tags == nil {
		fmt.Print("tags is nil")
	}
	fmt.Print(artical.Tags)
	err = db.Create(&artical).Error
	return
}

func GetPosts() (posts []interface{}, err error) {
	var post []Article
	db := sql.GetDB()
	// 使用 Preload 加载关联的 Tags
	err = db.Preload("Tags").Find(&post).Error
	if err != nil {
		return nil, err
	}
	for _, t := range post {
		tags := t.Tags
		// fmt.Print(tags)
		tagsreturn := []string{}
		for _, tag := range tags {
			// fmt.Print(tag.Name)
			tagsreturn = append(tagsreturn, tag.Name)
		}
		posts = append(posts, map[string]interface{}{
			"id":      t.ID,
			"title":   t.Title,
			"content": t.Content,
			"tags":    tagsreturn,
		})
	}
	return
}
