package models

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
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// Order represent the client purchase
type Order struct {
	ID         int `json:"id,omitempty`
	Status     int `json:"status,omitempty"`
	CustomerID int `json:"customer_id,omitempty"`
	ProductID  int `json:"product_id,omitemtpy"`
}
