package models

// Product represent the store
type Product struct {
	ID       int64   `json:"id, omitempty"`
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}
