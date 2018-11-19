package test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	return router
}

func PrepareRequest(router *gin.Engine, httpMethod string, url string, body io.Reader, handler gin.HandlerFunc) (*http.Request, *httptest.ResponseRecorder) {
	// Create a response recorder
	res := httptest.NewRecorder()

	// Define the route similar to its definition in the routes file
	router.Handle(httpMethod, url, handler)

	// Create a request to send to the above route
	req, _ := http.NewRequest(httpMethod, url, body)

	return req, res
}
