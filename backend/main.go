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
	test := r.Group("/test")
	{
		test.GET("/init", service.DBinit)
		test.POST("/create", service.CreatePost)
		test.GET("/getposts", service.GetPosts)
		test.DELETE("/delete/:uid", service.DeletePost)
		test.GET("/post/:uid", service.GetPost)
	}
	r.Run()
}
