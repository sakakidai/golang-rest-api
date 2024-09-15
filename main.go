package main

import (
	"golang-rest-api/config"
	"golang-rest-api/controllers"
	"golang-rest-api/db"
	"golang-rest-api/middleware"
	"golang-rest-api/repositories"
	"golang-rest-api/usecases"
	"golang-rest-api/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Config
	config.LoadConfig()

	// Logger
	utils.InitLogger()

	// Database
	db := db.ConnectDB()

	logger := utils.Logger()
	logger.Info("Go enviroment is " + os.Getenv("GO_ENV"))

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
	e.Use(middleware.Logger())
	e.Use(gin.Recovery())

	v1 := e.Group("api/v1")

	v1.GET("/welcome", welcomeController.Greet)

	u := v1.Group("/users")
	// u.POST("temporary_users", userController.TemporaryRegister)
	u.POST("/signup", userController.SignUp)
	u.POST("/login", userController.LogIn)
	u.POST("/confirm_email", userController.ConfirmEmail)

	ci := v1.Group("/content_items")
	ci.GET("/", contentItemController.GetAll)
	ci.POST("/", contentItemController.Create)

	logger.Info("Go server is running")
	e.Run(":8080")
}
