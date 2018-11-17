package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/quick-tournament-api-go/app/api"
	"github.com/sirupsen/logrus"
)

func InitializeRoutes(router *gin.Engine) {
	logrus.Info("Start initializing the api routes")

	router.GET("/ping", api.PongHandler)

}
