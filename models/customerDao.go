package models

import "database/sql"

// CustomerDao operates the database opereations
type CustomerDao struct {
	DB *sql.DB
}

// NewCustomerDao instanstiates a new CustomerDao
func NewCustomerDao(db *sql.DB) *CustomerDao {
	return &CustomerDao{
		DB: db,
	}
}

// Create inserts the new custome in the database
func (dao *CustomerDao) Create(customer Customer) (*Customer, error) {
	query := `
		INSERT INTO customers(name, email) VALUES($1, $2)
		RETURNING id
	`
	err := dao.DB.QueryRow(query, customer.Name, customer.Email).Scan(&customer.ID)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// Update changes the customer by id
func (dao *CustomerDao) Update(id string, customer Customer) error {
	stmt, err := dao.DB.Prepare("UPDATE customers SET name=$2, email=$3 WHERE id=$1")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, customer.Name, customer.Email)

	if err != nil {
		return err
	}

	return nil
}

// Read retrieves all customers
func (dao *CustomerDao) Read() ([]Customer, error) {
	var customers []Customer
	rows, err := dao.DB.Query("SELECT * FROM customers")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var name string
		var email string

		err := rows.Scan(&id, &name, &email)
		if err != nil {
			return nil, err
		}

		customers = append(customers, Customer{
			ID:    id,
			Name:  name,
			Email: email,
		})
	}

	return customers, nil
}

// Remove excluses the customer by id
func (dao CustomerDao) Remove(id string) (err error) {
	query := `DELETE FROM customers WHERE id=$1`
	stmt, err := dao.DB.Prepare(query)
	_, err = stmt.Exec(id)
	return
}
