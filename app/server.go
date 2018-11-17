package app

import (
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("$PORT must be set")
	}

	logrus.Info("Server is listening on http://localhost:" + port)
	router.Run(":" + port)
}
