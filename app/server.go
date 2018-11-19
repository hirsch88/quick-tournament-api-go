package app

import (
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/quick-tournament-api-go/app/api/middlewares"
	"github.com/hirsch88/quick-tournament-api-go/app/lib"
	"github.com/sirupsen/logrus"
	"github.com/thinkerou/favicon"
)

func Bootstrap() {
	logrus.Info("Start bootstraping the api")

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(lib.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(middlewares.CORSMiddleware())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(helmet.Default())
	router.Use(helmet.NoCache())
	router.Use(favicon.New("./static/favicon.ico"))
	router.Use(static.Serve("/", static.LocalFile("./static", false)))

	// Defines the api endpoints and constructs the manuel ioc
	InitializeRoutes(router)

	// Reads the port value from the env-values
	port := os.Getenv("PORT")

	// Runs server on the given port number
	logrus.Info("Server is listening on http://localhost:" + port)
	router.Run(":" + port)
}
