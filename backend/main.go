package main

import (
	"goblogeasyg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", service.Home)
	r.Run()
}
