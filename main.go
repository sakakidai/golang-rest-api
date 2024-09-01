package main

import (
	"golang-rest-api/controllers"
	"golang-rest-api/db"
	"golang-rest-api/repositories"
	"golang-rest-api/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.NewDB()
	wc := controllers.NewWelcomeController()
	contentItemRepository := repositories.NewContentItemRepository(db)
	contentItemUsecase := usecases.NewContentItemUsecase(contentItemRepository)
	contentItemController := controllers.NewContentItemController(contentItemUsecase)

	r := gin.Default()
	v1 := r.Group("api/v1")

	// welcome
	v1.GET("/welcome", wc.Greet)

	// content_items
	ci := v1.Group("content_items")
	ci.GET("/", contentItemController.GetAll)
	ci.POST("/", contentItemController.Create)

	r.Run(":8080")
}
