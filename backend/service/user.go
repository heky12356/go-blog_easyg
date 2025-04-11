package service

import (
	"net/http"

	sql "goblogeasyg/sql/model"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Login
func Login(c *gin.Context) {
	var data User
	var user sql.User
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.Username == "" || data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password cannot be empty"})
		return
	}
	user, err := sql.GetUserByUsername(data.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Password != data.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is wrong"})
		return
	}
	c.JSON(200, gin.H{"message": "Login"})
}

// Register
func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := sql.CreateUser(sql.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Register"})
}
