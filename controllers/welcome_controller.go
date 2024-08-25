package controllers

import (
	"github.com/gin-gonic/gin"
)

type IWelcomeController interface {
	Greet(c *gin.Context)
	// GetUserById(c *gin.Context)
}

type WelcomeController struct{}

type Welcome struct {
	Greet    string `json:"greet"`
	Doc      string `json:"link_to_doc"`
	Github   string `json:"github"`
	Examples string `json:"examples"`
}

func NewWelcomeController() IWelcomeController {
	return &WelcomeController{}
}

func (uc *WelcomeController) Greet(c *gin.Context) {
	welcome := Welcome{
		Greet:    "Welcome to letsGo",
		Doc:      "https://letsgo-framework.github.io/",
		Github:   "https://github.com/letsgo-framework/letsgo",
		Examples: "Coming Soon",
	}
	c.JSON(200, welcome)
	c.Done()
}
