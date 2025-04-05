package sql

import (
	"goblogeasyg/sql"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name    string    `gorm:"uniqueIndex"`
	Aticles []Article `gorm:"many2many:article_tags;"`
}
type Article struct {
	BaceModel
	Content string
	Title   string
	Uid     string `gorm:"primarykey"`
	Tags    []Tag  `gorm:"many2many:article_tags;"`
}

func AutoMigrate() (err error) {
	db := sql.GetDB()
	err = db.AutoMigrate(&Article{}, &Tag{})
	return
}

func CreatePost(artical Article) error {
	db := sql.GetDB()
	return db.Transaction(func(tx *gorm.DB) error {
		// 处理每一个tag
		for i, tag := range artical.Tags {
			var exitTag Tag
			if err := tx.Where("name = ?", tag.Name).First(&exitTag).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					err = tx.Create(&tag).Error
					if err != nil {
						return err
					}
					artical.Tags[i] = tag
				}
			} else {
				artical.Tags[i] = exitTag
			}
		}

		// 创建文章
		err := tx.Create(&artical).Error
		if err != nil {
			return err
		}

		return nil
	})
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
			"title":   t.Title,
			"content": t.Content,
			"uid":     t.Uid,
			"tags":    tagsreturn,
		})
	}
	return
}

func DeletePost(uid string) (err error) {
	var artical Article
	db := sql.GetDB()
	err = db.Where("uid = ?", uid).First(&artical).Error
	if err != nil {
		return err
	}
	err = db.Model(&artical).Association("Tags").Clear()
	if err != nil {
		return err
	}
	err = db.Model(&artical).Delete(&artical).Error
	if err != nil {
		return err
	}
	return
}
