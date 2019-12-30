package models

import "database/sql"

// Product represent the store
type Product struct {
	ID       int64         `json:"id, omitempty"`
	Name     string        `json:"name,omitempty"`
	Price    float64       `json:"price,omitempty"`
	Quantity int           `json:"quantity,omitempty"`
	UserID   sql.NullInt64 `json:"user_id,omitempty"`
}
