package models

import "database/sql"

// ProductDao operates basic commom operation in products tables
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
	query := `
		INSERT INTO products 
		(name, price, quantity, userid) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := pd.Conn.QueryRow(query, product.Name, product.Price, product.Quantity, product.UserID).Scan(&product.ID)
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

// Update changes the product by id
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

	rows, err := pd.Conn.Query("SELECT id, name, price, quantity, userid FROM products ORDER BY id")
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var id int64
		var name string
		var price float64
		var quantity int
		var userID sql.NullInt64

		err := rows.Scan(&id, &name, &price, &quantity, &userID)

		if err != nil {
			return nil, err
		}

		products = append(products, Product{
			ID:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
			UserID:   userID,
		})
	}

	return products, nil
}

// Find retrieves a unique product instance
func (pd *ProductDao) Find(id int64) (*Product, error) {
	var product Product
	row := pd.Conn.QueryRow("SELECT * FROM products WHERE id=$1", id)
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.UserID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
