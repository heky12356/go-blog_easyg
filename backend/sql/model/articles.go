package sql

import (
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
