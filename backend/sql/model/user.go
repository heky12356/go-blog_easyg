package sql

import (
	"goblogeasyg/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func AutoMigrateUser() (err error) {
	db := sql.GetDB()
	err = db.AutoMigrate(&User{})
	return
}

func CreateUser(user User) error {
	db := sql.GetDB()
	return db.Create(&user).Error
}

func GetUserByEmail(email string) (user User, err error) {
	db := sql.GetDB()
	err = db.Where("email = ?", email).First(&user).Error
	return
}

func GetUserByUsername(username string) (user User, err error) {
	db := sql.GetDB()
	err = db.Where("username = ?", username).First(&user).Error
	return
}

func DeleteUser(user User) (err error) {
	db := sql.GetDB()
	err = db.Delete(&user).Error
	return
}
