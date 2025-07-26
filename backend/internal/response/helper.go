package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 成功响应
func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: message,
		Data:    data,
		Success: true,
	})
}

// 错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    nil,
		Success: false,
	})
}
