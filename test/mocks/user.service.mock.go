package mocks

import (
	"github.com/hirsch88/quick-tournament-api-go/app/api/dtos"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) GetAll() []dtos.User {
	args := m.Called()
	return args.Get(0).([]dtos.User)
}

func (m *UserServiceMock) GetByID(id int64) (dtos.User, error) {
	args := m.Called(id)
	return args.Get(0).(dtos.User), args.Error(1)
}

func (m *UserServiceMock) Create(newUser dtos.User) (dtos.User, error) {
	args := m.Called(newUser)
	return args.Get(0).(dtos.User), args.Error(1)
}

func (m *UserServiceMock) DeleteByID(id int64) (dtos.User, error) {
	args := m.Called(id)
	return args.Get(0).(dtos.User), args.Error(1)
}
