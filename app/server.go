package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Bootstrap() {
	logrus.Info("Start bootstraping the api")

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	InitializeRoutes(router)

	gin.Logger()

	logrus.Info("Server is listening on http://localhost:8080")
	router.Run(":8080")
}
