package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/quick-tournament-api-go/app/api/controllers"
	"github.com/hirsch88/quick-tournament-api-go/app/api/repositories"
	"github.com/hirsch88/quick-tournament-api-go/app/api/services"
	"github.com/hirsch88/quick-tournament-api-go/app/database"
	"github.com/sirupsen/logrus"
)

func InitializeRoutes(router *gin.Engine) {
	logrus.Info("Start initializing the api routes")

	api := router.Group("/api")

	// // Create the repositories
	baseRepository := repositories.NewBaseRepository(database.GetDB())
	userRepository := repositories.NewUserRepository(baseRepository)

	// // Create the services
	userService := services.NewUserService(userRepository)

	// // Create controllers
	userController := controllers.NewUserController(userService)
	apiController := controllers.NewAPIController()

	// Register controllers to the api path
	// /api
	api.GET("/", apiController.GetInfo)
	// /users
	api.Group("/users").
		GET("", userController.GetAll)

}
