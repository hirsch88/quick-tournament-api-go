package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestPongHandler(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Define the route similar to its definition in the routes file
	router.GET("/ping", PongHandler)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Create the service and process the above request.
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Convert the JSON response to a map
	var body map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &body)
	assert.Nil(t, err)
	assert.Equal(t, body["message"], "pong")
}
