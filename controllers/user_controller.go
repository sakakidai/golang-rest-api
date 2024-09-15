package controllers

import (
	"golang-rest-api/config"
	"golang-rest-api/models"
	"golang-rest-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

type IUserController interface {
	// TemporaryRegister(c *gin.Context)
	SignUp(c *gin.Context)
}

type userController struct {
	uu usecases.IUserUsecase
}

func NewUserController(uu usecases.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindWith(&user, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	token, err := uc.uu.SignUp(c, user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error1"})
		return
	}

	if config.GetConfig().EmailVerification.Enabled {
		if len(token) == 0 {
			c.Error(err).SetType(gin.ErrorTypePrivate)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error2"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"confirmToken": token})
		return
	}

	c.Status(http.StatusOK)
}

// Private method

func logger(c *gin.Context) *zap.Logger {
	loggerInterface, exists := c.Get("logger")
	if !exists {
		panic("logger not found in context")
	}

	return loggerInterface.(*zap.Logger)
}
