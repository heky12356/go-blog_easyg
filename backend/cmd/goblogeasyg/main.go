package main

import (
	"goblogeasyg/internal/middleware"
	"goblogeasyg/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", service.Home)

	apiuser := r.Group("/api/user")
	{
		apiuser.POST("/register", service.Register)
		apiuser.POST("/login", service.Login)
		apiuser.POST("/logout", middleware.Auth(), service.Logout)
		apiuser.GET("/getalluser", service.GetallUser)
		apiuser.POST("/refreshaccesstoken", service.RefreshAccessToken)
		apiuser.POST("/verify", middleware.Auth())
	}
	postapi := r.Group("/api/post")
	{
		postapi.POST("/create", middleware.Auth(), service.CreatePost)
		postapi.DELETE("/delete/:uid", middleware.Auth(), service.DeletePost)

		postapi.GET("/getposts", service.GetPosts)
		postapi.GET("/post/:uid", service.GetPost)
	}
	r.Run()
}
