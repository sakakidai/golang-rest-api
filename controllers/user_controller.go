package controllers

import (
	"golang-rest-api/models"
	"golang-rest-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

// func (uc *userController) TemporaryRegister(c *gin.Context) {
// 	c.JSON(http.StatusOK, 'ssss')
// }

func (uc *userController) SignUp(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindWith(&user, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	res, err := uc.uu.SignUp(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, res)
}
