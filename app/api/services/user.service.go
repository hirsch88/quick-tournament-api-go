package services

import (
	"github.com/hirsch88/quick-tournament-api-go/app/api/dtos"
	"github.com/hirsch88/quick-tournament-api-go/app/api/models"
	"github.com/hirsch88/quick-tournament-api-go/app/api/repositories"
	funk "github.com/thoas/go-funk"
)

// UserService ...
type UserService interface {
	GetAll() []dtos.User
	GetByID(id int64) (dtos.User, error)
	Create(newUser dtos.User) (dtos.User, error)
	DeleteByID(id int64) (dtos.User, error)
	// GetByID(id int64) (dtos.User, bool)
	// DeleteByID(id int64) bool
	// UpdatePosterAndGenreByID(id int64, poster string, genre string) (dtos.User, error)
}

// NewUserService returns the default User service.
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.UserRepository
}

// GetAll returns all users.
func (s *userService) GetAll() []dtos.User {
	var (
		users    []models.User
		userDtos []dtos.User
	)
	users = s.repo.FindAll()

	userDtos = funk.Map(users, func(user models.User) dtos.User {
		return user.ToJSON()
	}).([]dtos.User)

	return userDtos
}

// GetByID returns a user based on its id.
func (s *userService) GetByID(id int64) (dtos.User, error) {
	var (
		user models.User
		err  error
	)
	user, err = s.repo.FindByID(id)
	return user.ToJSON(), err
}

// Create ...
func (s *userService) Create(newUser dtos.User) (dtos.User, error) {
	var (
		user models.User
		err  error
	)
	user = models.NewUserFromJSON(newUser)
	user, err = s.repo.Create(user)

	return user.ToJSON(), err
}

// Delete ...
func (s *userService) DeleteByID(id int64) (dtos.User, error) {
	var (
		user models.User
		err  error
	)
	user, err = s.repo.FindByID(id)

	if err != nil {
		return user.ToJSON(), err
	}

	user, err = s.repo.Delete(user)

	return user.ToJSON(), err
}
