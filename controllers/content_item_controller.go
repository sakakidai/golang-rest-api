package controllers

import (
	"golang-rest-api/models"
	"golang-rest-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type IContentItemController interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
}

type contentItemController struct {
	ciu usecases.IContentItemUsecase
}

func NewContentItemController(ciu usecases.IContentItemUsecase) IContentItemController {
	return &contentItemController{ciu}
}

func (cic *contentItemController) GetAll(c *gin.Context) {
	res, err := cic.ciu.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}

func (cic *contentItemController) Create(c *gin.Context) {
	ci := models.ContentItem{}
	err := c.ShouldBindWith(&ci, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	res, err := cic.ciu.Create(ci)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}
