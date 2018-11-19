package controllers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hirsch88/quick-tournament-api-go/test/mocks"

	"github.com/hirsch88/quick-tournament-api-go/app/api/dtos"
	"github.com/hirsch88/quick-tournament-api-go/test"

	"github.com/stretchr/testify/assert"
)

func TestUserGetAll(t *testing.T) {
	userServiceMock := new(mocks.UserServiceMock)
	controller := NewUserController(userServiceMock)

	newUserDto := dtos.User{ID: 1, FirstName: "Hans", LastName: "Muster"}
	newUserDtos := []dtos.User{newUserDto}
	userServiceMock.On("GetAll").Return(newUserDtos)

	router := test.GetRouter()
	req, res := test.PrepareRequest(router, "GET", "/api", nil, controller.GetAll)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	userBody := []dtos.User{}
	json.NewDecoder(res.Body).Decode(&userBody)
	assert.Equal(t, "Hans", userBody[0].FirstName)
	assert.Equal(t, "Muster", userBody[0].LastName)

}
