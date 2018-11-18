package models

import "github.com/hirsch88/quick-tournament-api-go/app/api/dtos"

// User ...
type User struct {
	BaseModel

	FirstName string `gorm:"not null; type:varchar(100)"`
	LastName  string `gorm:"not null; type:varchar(100)"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}

// NewUserFromJSON ...
func NewUserFromJSON(userDto dtos.User) User {
	return User{
		BaseModel: BaseModel{
			ID: userDto.ID,
		},
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
	}
}

// ToJSON ...
func (u *User) ToJSON() dtos.User {
	return dtos.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
