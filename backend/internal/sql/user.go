package sql

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func AutoMigrateUser() (err error) {
	err = db.AutoMigrate(&User{})
	return
}

func CreateUser(user User) error {
	return db.Create(&user).Error
}

func GetUserByEmail(email string) (user User, err error) {
	err = db.Where("email = ?", email).First(&user).Error
	return
}

func GetUserByUsername(username string) (user User, err error) {
	err = db.Where("username = ?", username).First(&user).Error
	return
}

func DeleteUser(user User) (err error) {
	err = db.Delete(&user).Error
	return
}

func GetallUser() (data []map[string]interface{}, err error) {
	var user []User
	err = db.Find(&user).Error
	for _, u := range user {
		data = append(data, map[string]interface{}{
			"username": u.Username,
			"email":    u.Email,
		})
	}
	return
}
