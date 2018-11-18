package dtos

// User ...
type User struct {
	ID        uint64 `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty" validate:"required"`
	LastName  string `json:"lastname,omitempty" validate:"required"`
}
