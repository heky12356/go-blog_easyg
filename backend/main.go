package main

import (
	"log"
	"os"

	"goblogeasyg/service"
	"goblogeasyg/sql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DbName := os.Getenv("DB_NAME")
	err = sql.Init(DbName)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.GET("/", service.Home)
	r.GET("/init", service.DBinit)
	postapi := r.Group("/api/post")
	{
		postapi.POST("/create", service.CreatePost)
		postapi.GET("/getposts", service.GetPosts)
		postapi.DELETE("/delete/:uid", service.DeletePost)
		postapi.GET("/post/:uid", service.GetPost)
	}
	r.Run()
}
