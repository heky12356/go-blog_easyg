package main

import (
	"goblogeasyg/internal/handler"
	"goblogeasyg/internal/middleware"
	"goblogeasyg/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	postService := service.NewPostService()
	postHandler := handler.NewPostHandler(postService)

	r := gin.Default()
	r.GET("/", service.Home)

	apiuser := r.Group("/api/user")
	{
		apiuser.POST("/register", userHandler.Register)
		apiuser.POST("/login", userHandler.Login)
		apiuser.POST("/logout", middleware.Auth(), userHandler.Logout)
		apiuser.GET("/getalluser", userHandler.GetallUser)
		apiuser.POST("/refreshaccesstoken", userHandler.RefreshAccessToken)
		apiuser.POST("/verify", middleware.Auth())
	}
	postapi := r.Group("/api/post")
	{
		postapi.POST("/create", middleware.Auth(), postHandler.CreatePost)
		postapi.DELETE("/delete/:uid", middleware.Auth(), postHandler.DeletePost)
		postapi.GET("/getposts", postHandler.GetAllPost)
		postapi.GET("/post/:uid", postHandler.GetPost)
	}
	r.Run(":8080")
}
