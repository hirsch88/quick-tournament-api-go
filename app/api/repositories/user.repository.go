package repositories

import "github.com/hirsch88/quick-tournament-api-go/app/api/models"

// UserRepository ...
type UserRepository interface {
	FindAll() []models.User
	FindByID(id int64) (models.User, error)
	Create(user models.User) (models.User, error)
	Delete(user models.User) (models.User, error)
}

// NewUserRepository ...
func NewUserRepository(base BaseRepository) UserRepository {
	return &userRepository{
		base: base,
	}
}

type userRepository struct {
	base BaseRepository
}

// FindAll ...
func (r *userRepository) FindAll() []models.User {
	var users []models.User
	r.base.FindAll(&users)
	return users
}

// FindByID
func (r *userRepository) FindByID(id int64) (models.User, error) {
	var (
		user models.User
		err  error
	)
	err = r.base.FindOneByID(&user, id)
	return user, err
}

// Create
func (r *userRepository) Create(user models.User) (models.User, error) {
	var err error
	err = r.base.Create(&user)
	return user, err
}

// Delete
func (r *userRepository) Delete(user models.User) (models.User, error) {
	var err error
	err = r.base.Delete(&user)
	return user, err
}
