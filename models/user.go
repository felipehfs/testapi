package models

import "github.com/go-playground/validator"

// User represents the authenticated user
type User struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

// IsValid checks if each field is correct
func (user User) IsValid() error {
	validate := validator.New()
	return validate.Struct(user)
}
