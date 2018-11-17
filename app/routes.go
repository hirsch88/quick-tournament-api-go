package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"w3tec.ch/go-heroku/app/api"
)

func InitializeRoutes(router *gin.Engine) {
	logrus.Info("Start initializing the api routes")

	router.GET("/ping", api.PongHandler)

}
