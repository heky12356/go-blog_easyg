package service

import (
	"errors"
	"fmt"
	"time"

	"goblogeasyg/internal/cache"
	sql "goblogeasyg/internal/sql"
	utils "goblogeasyg/internal/utils/jwt"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Register(username string, password string, confirmPassword string, email string) error
	Login(username string, password string) (token []map[string]string, err error)
	GetallUser() (users []User, err error)
	RefreshAccessToken(refreshToken string) (token string, err error)
	Logout(token string) error
}

type UserService struct{}

func NewUserService() UserServiceInterface {
	return &UserService{}
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// refresh access token
func (u *UserService) RefreshAccessToken(refreshToken string) (token string, err error) {
	token, err = utils.RefreshToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("refresh token failed: %w", err)
	}
	return token, nil
}

// get all user
func (u *UserService) GetallUser() (users []User, err error) {
	data, err := sql.GetallUser()
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		users = append(users, User{
			Username: d["username"].(string),
			Email:    d["email"].(string),
		})
	}
	return users, nil
}

// Login
func (u *UserService) Login(username string, password string) (token []map[string]string, err error) {
	var user sql.User
	if username == "" || password == "" {
		return nil, errors.New("username or password cannot be empty")
	}
	user, err = sql.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("get user by username failed: %w", err)
	}

	// Compare the hashed password with the provided password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("password is wrong")
	}

	refreshtoken, err := utils.CreateRefreshToken(user.Username)
	if err != nil {
		return nil, fmt.Errorf("create refresh token failed: %w", err)
	}
	acessstoken, err := utils.CreateAssessToken(user.Username)
	if err != nil {
		return nil, fmt.Errorf("create access token failed: %w", err)
	}
	token = []map[string]string{
		{
			"refreshToken": refreshtoken,
			"accessToken":  acessstoken,
		},
	}
	return token, nil
}

// Register
func (u *UserService) Register(username string, password string, confirmPassword string, email string) error {
	if username == "" || password == "" || confirmPassword == "" || email == "" {
		return errors.New("username or password or confirm password or email cannot be empty")
	}
	if password != confirmPassword {
		return errors.New("password and confirm password not match")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err)
		return fmt.Errorf("hash password failed: %w", err)
	}

	err = sql.CreateUser(sql.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	})
	if err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}

// Logout handles user logout by invalidating the current token
func (u *UserService) Logout(token string) error {
	// 解析token以获取过期时间
	claims, err := utils.ParseToken(token)
	if err != nil {
		return fmt.Errorf("parse token failed: %w", err)
	}

	// 计算token剩余有效期
	expiration := time.Until(time.Unix(claims.ExpiresAt, 0))
	if expiration <= 0 {
		return fmt.Errorf("token expired")
	}

	// 将token加入黑名单，使用剩余有效期作为过期时间
	err = cache.AddToBlacklist(token, expiration)
	if err != nil {
		return fmt.Errorf("add token to blacklist failed: %w", err)
	}

	return nil
}
