package service

import (
	"net/http"

	sql "goblogeasyg/sql/model"
	"goblogeasyg/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// refresh access token
func RefreshToken(c *gin.Context) {
	var data User
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.Username == "" || data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password cannot be empty"})
		return
	}
	token, err := utils.CreateAssessToken(data.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

// get all user
func GetallUser(c *gin.Context) {
	data, err := sql.GetallUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"users": data})
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

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is wrong"})
		return
	}

	refreshtoken, err := utils.CreateRefreshToken(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	acessstoken, err := utils.CreateAssessToken(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Login", "refresh_token": refreshtoken, "access_token": acessstoken})
}

// Register
func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "host error"})
		return
	}

	err = sql.CreateUser(sql.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Email:    user.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Register"})
}
