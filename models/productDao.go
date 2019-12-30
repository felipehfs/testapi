package models

import "database/sql"

// ProductDao
type ProductDao struct {
	Conn *sql.DB
}

// NewProductDao creates the productDao instance
func NewProductDao(db *sql.DB) *ProductDao {
	return &ProductDao{
		Conn: db,
	}
}

// Create the new Product
func (pd *ProductDao) Create(product Product) (*Product, error) {
	smtp, err := pd.Conn.Prepare("INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3)")
	if err != nil {
		return nil, err
	}

	_, err = smtp.Exec(product.Name, product.Price, product.Quantity)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Remove makes the drop of the database
func (pd *ProductDao) Remove(id int64) error {
	stmt, err := pd.Conn.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(&id)
	return err
}

// Update
func (pd *ProductDao) Update(product Product, id string) error {
	stmt, err := pd.Conn.Prepare("UPDATE products SET name=$2, price=$3, quantity=$4 WHERE id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, product.Name, product.Price, product.Quantity)
	if err != nil {
		return err
	}
	return nil
}

// Read
func (pd *ProductDao) Read() ([]Product, error) {
	var products []Product

	rows, err := pd.Conn.Query("SELECT id, name, price, quantity FROM products ORDER BY id")
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var id int64
		var name string
		var price float64
		var quantity int

		err := rows.Scan(&id, &name, &price, &quantity)

		if err != nil {
			return nil, err
		}

		products = append(products, Product{
			ID:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
		})
	}

	return products, nil
}

// Find retrieves a unique product instance
func (pd *ProductDao) Find(id int64) (*Product, error) {
	var product Product
	row := pd.Conn.QueryRow("SELECT * FROM products WHERE id=$1", id)
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
