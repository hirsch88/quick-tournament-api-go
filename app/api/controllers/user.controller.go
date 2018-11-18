package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/quick-tournament-api-go/app/api/services"
)

// UserController ...
type UserController interface {
	GetAll(c *gin.Context)
}

// NewUserController ...
func NewUserController(service services.UserService) UserController {
	return &userController{
		service: service,
	}
}

type userController struct {
	service services.UserService
}

// GetAll ...
func (c *userController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, c.service.GetAll())
}
