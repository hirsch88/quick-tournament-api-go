package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

// APIController ...
type APIController interface {
	GetInfo(c *gin.Context)
}

// NewAPIController ...
func NewAPIController() APIController {
	return &apiController{}
}

type apiController struct{}

// GetAll ...
func (c *apiController) GetInfo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"title":   os.Getenv("API_TITLE"),
		"version": os.Getenv("API_VERSION"),
	})
}
