package models

import "github.com/go-playground/validator"

const (
	// OrderSuccessfully is order completed
	OrderSuccessfully = iota
	// OrderCancelled is order rejected
	OrderCancelled
	// OrderProccessing is order waiting status
	OrderProccessing
)

// Customer represents the client
type Customer struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
}

// IsValid checks if the order is correct
func (customer Customer) IsValid() error {
	validate := validator.New()
	return validate.Struct(customer)
}

// Order represent the client purchase
type Order struct {
	ID         int   `json:"id,omitempty`
	Status     int   `json:"status,omitempty"`
	CustomerID int   `json:"customer_id,omitempty" validate:"required"`
	ProductID  int   `json:"product_id,omitempty" validate:"required"`
	Author     int64 `json:"author"`
}

// IsValid checks if the order is correct
func (order Order) IsValid() error {
	validate := validator.New()
	return validate.Struct(order)
}
