package main

import (
	"golang-rest-api/config"
	"golang-rest-api/controllers"
	"golang-rest-api/db"
	"golang-rest-api/middleware"
	"golang-rest-api/repositories"
	"golang-rest-api/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	// Config
	config.LoadConfig()

	// Database
	db := db.ConnectDB()

	// Repository
	var (
		userRepository        = repositories.NewUserRepository(db)
		contentItemRepository = repositories.NewContentItemRepository(db)
	)

	// Usecase
	var (
		userUsecase        = usecases.NewUserUsecase(userRepository)
		contentItemUsecase = usecases.NewContentItemUsecase(contentItemRepository)
	)

	// Controller
	var (
		welcomeController     = controllers.NewWelcomeController()
		userController        = controllers.NewUserController(userUsecase)
		contentItemController = controllers.NewContentItemController(contentItemUsecase)
	)

	// Router
	e := gin.Default()
	e.Use(middleware.ZapLogger())
	e.Use(gin.Recovery())

	v1 := e.Group("api/v1")

	v1.GET("/welcome", welcomeController.Greet)

	u := v1.Group("/users")
	// u.POST("temporary_users", userController.TemporaryRegister)
	u.POST("/signup", userController.SignUp)

	ci := v1.Group("/content_items")
	ci.GET("/", contentItemController.GetAll)
	ci.POST("/", contentItemController.Create)

	e.Run(":8080")
}
