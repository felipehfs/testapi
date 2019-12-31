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

// Find searches the Order and shows a complete information
func (dao *OrderDao) Find(id string) (map[string]interface{}, error) {
	query := `
		SELECT 
			o.id, o.status, o.author,
			p.name, p.price,
			c.id, c.name, c.email
		FROM 
			orders o 
		INNER JOIN products p ON o.productid = p.id
		INNER JOIN customers c ON o.customerid = c.id
		WHERE o.id = $1
	`
	var orderID int
	var status int
	var orderAuthor int
	var name string
	var price float64
	var customerID int
	var customerName string
	var customerEmail string

	order := make(map[string]interface{})

	err := dao.DB.QueryRow(query, id).Scan(&orderID, &status, &orderAuthor,
		&name, &price, &customerID, &customerName, &customerEmail)
	if err != nil {
		return nil, err
	}

	order["product"] = map[string]interface{}{
		"name":  name,
		"price": price,
	}

	order["order"] = map[string]interface{}{
		"id":     orderID,
		"status": status,
	}

	order["customer"] = map[string]interface{}{
		"id":    customerID,
		"name":  customerName,
		"email": customerEmail,
	}
	return order, nil
}

// Update changes the orders
func (dao *OrderDao) Update(order Order, id string) (*Order, error) {
	query := `
		UPDATE orders SET status=$2, customerid=$3, productid=$4
		WHERE id=$1 
	`
	stmt, err := dao.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(order.ID, order.Status, order.CustomerID, order.ProductID)
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
