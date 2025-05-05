package main

import (
	"log"
	"os"

	"goblogeasyg/internal/cache"
	"goblogeasyg/internal/middleware"
	"goblogeasyg/internal/service"
	"goblogeasyg/internal/sql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DbName := os.Getenv("DB_NAME")
	err = sql.Init(DbName)
	if err != nil {
		log.Fatal(err)
	}

	// 初始化Redis
	err = cache.InitRedis()
	if err != nil {
		log.Fatal("Redis初始化失败:", err)
	}
	r := gin.Default()
	r.GET("/", service.Home)
	r.GET("/init", service.DBinit)

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
