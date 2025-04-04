package sql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(dbname string) (err error) {
	db, err = gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	return
}

func GetDB() *gorm.DB {
	return db
}
