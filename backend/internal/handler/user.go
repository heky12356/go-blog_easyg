package handler

import (
	"fmt"
	"strings"

	"goblogeasyg/internal/response"
	"goblogeasyg/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetallUser(c *gin.Context)
	RefreshAccessToken(c *gin.Context)
	Logout(c *gin.Context)
}
type UserHandler struct {
	userService service.UserServiceInterface
}

// GetallUser implements UserHandlerInterface.
func (u *UserHandler) GetallUser(c *gin.Context) {
	users, err := u.userService.GetallUser()
	if err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, users, "get all user success")
}

// Login implements UserHandlerInterface.
func (u *UserHandler) Login(c *gin.Context) {
	var data LoginRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, err.Error())
		return
	}
	token, err := u.userService.Login(data.Username, data.Password)
	if err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, token, "login success")
}

// RefreshAccessToken implements UserHandlerInterface.
func (u *UserHandler) RefreshAccessToken(c *gin.Context) {
	var data RefreshAccessTokenRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, err.Error())
		return
	}
	token, err := u.userService.RefreshAccessToken(data.RefreshToken)
	if err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("refresh access token failed: %s", err.Error()))
		return
	}
	response.SuccessResponse(c, token, "refresh access token success")
}

// Register implements UserHandlerInterface.
func (u *UserHandler) Register(c *gin.Context) {
	var data RegisterRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, err.Error())
		return
	}
	if err := u.userService.Register(data.Username, data.Password, data.ConfirmPassword, data.Email); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, nil, "register success")
}

// Logout implements UserHandlerInterface.
func (u *UserHandler) Logout(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		response.ErrorResponse(c, response.CodeUnauthorized, "未提供认证令牌")

		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	if err := u.userService.Logout(tokenString); err != nil {
		response.ErrorResponse(c, response.CodeBadRequest, fmt.Sprintf("注销失败: %s", err.Error()))
		return
	}
	response.SuccessResponse(c, nil, "注销成功")
}

func NewUserHandler(userService service.UserServiceInterface) UserHandlerInterface {
	return &UserHandler{userService: userService}
}
