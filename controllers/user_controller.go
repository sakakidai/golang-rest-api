package controllers

import (
	"golang-rest-api/models"
	"golang-rest-api/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IUserController interface {
	// TemporaryRegister(c *gin.Context)
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
	CreateConfirmToken(c *gin.Context)
	ConfirmEmail(c *gin.Context)
}

type userController struct {
	uu usecases.IUserUsecase
}

func NewUserController(uu usecases.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c *gin.Context) {
	// ログインしてないことをチェック

	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	if err := uc.uu.SignUp(c, user); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.Status(http.StatusOK)
}

func (uc *userController) LogIn(c *gin.Context) {
	// ログインしてないことをチェック

	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	tokenString, err := uc.uu.LogIn(c, user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "The email address or password is incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authenticationToken": tokenString})
	return
}

func (uc *userController) ConfirmEmail(c *gin.Context) {
	// ログインしていることをチェックする

	var requestBody struct {
		ConfirmToken string `json:"confirm_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	if err := uc.uu.ConfirmEmail(c, requestBody.ConfirmToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (uc *userController) CreateConfirmToken(c *gin.Context) {
	// ログインしていることをチェックする

	var userID uint
	var err error

	userIDStr := c.GetHeader("X-User-ID")
	userIDfloat, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid X-User-ID"})
		return
	}
	userID = uint(userIDfloat)

	err = uc.uu.CreateConfrimToken(c, &userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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
