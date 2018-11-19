package controllers

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/hirsch88/quick-tournament-api-go/app/api/dtos"

	"github.com/hirsch88/quick-tournament-api-go/test"

	"github.com/stretchr/testify/assert"
)

func TestAPIGetAPIInfo(t *testing.T) {
	os.Setenv("API_TITLE", "Title")
	os.Setenv("API_VERSION", "0.0.0")

	apiController := NewAPIController()

	router := test.GetRouter()
	req, res := test.PrepareRequest(router, "GET", "/api", nil, apiController.GetInfo)
	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)

	// Convert the JSON response to a map
	apiBody := dtos.API{}
	json.NewDecoder(res.Body).Decode(&apiBody)
	assert.NotEmpty(t, "Title", apiBody.Name)
	assert.NotEmpty(t, "0.0.0", apiBody.Version)

}
