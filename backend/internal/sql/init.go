package sql

import (
	"fmt"

	"goblogeasyg/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(config.GetConfig().DateBase.Litedb), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}
	err = AutoMigrateArticle()
	if err != nil {
		panic(fmt.Errorf("failed to migrate Article: %v", err))
	}
	err = AutoMigrateUser()
	if err != nil {
		panic(fmt.Errorf("failed to migrate User: %v", err))
	}
}
