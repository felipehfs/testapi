package models

import "database/sql"

// OrderDao operates commom database operation in orders table
type OrderDao struct {
	DB *sql.DB
}

// NewOrderDao instantiates the OrderDao
func NewOrderDao(db *sql.DB) *OrderDao {
	return &OrderDao{
		DB: db,
	}
}

// Create inserts the new data in the orders table
func (dao *OrderDao) Create(order Order) (*Order, error) {
	query := `
		INSERT INTO orders (status, customerid, productid, author)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	order.Status = OrderProccessing
	err := dao.DB.QueryRow(query, order.Status, order.CustomerID, order.ProductID, order.Author).Scan(&order.ID)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// Read retrieves all orders created by logged user
func (dao *OrderDao) Read(author int64) ([]Order, error) {

	var orders []Order
	rows, err := dao.DB.Query("SELECT * FROM orders WHERE author = $1", author)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order Order
		rows.Scan(
			&order.ID,
			&order.Status,
			&order.CustomerID,
			&order.ProductID,
			&order.Author)

		orders = append(orders, order)
	}

	return orders, nil
}
