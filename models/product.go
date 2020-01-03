package models

import (
	"database/sql"

	"github.com/go-playground/validator"
)

// Product represent the store
type Product struct {
	ID       int64         `json:"id,omitempty"`
	Name     string        `json:"name,omitempty" validate:"required"`
	Price    float64       `json:"price,omitempty" validate:"required,gte=0"`
	Quantity int           `json:"quantity,omitempty" validate:"gte=0"`
	UserID   sql.NullInt64 `json:"user_id,omitempty"`
}

// IsValid checks if the product is correct
func (p Product) IsValid() error {
	validate := validator.New()
	return validate.Struct(p)
}
