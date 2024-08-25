package main

import (
	"golang-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")

	wc := controllers.NewWelcomeController()
	v1.GET("/welcome", wc.Greet)

	r.Run(":8080")
}
