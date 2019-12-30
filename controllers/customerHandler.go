package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/felipehfs/testapi/models"
)

// CreateCustomer controller operates the insert in the database
func CreateCustomer(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customerDao := models.NewCustomerDao(db)
		var customer models.Customer

		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			http.Error(w, "Erro na sintaxe do JSON", http.StatusBadRequest)
			return
		}

		if err := customerDao.Create(customer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(customer)
	})
}

// ReadCustomer retrieves customers
func ReadCustomer(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customerDao := models.NewCustomerDao(db)
		customers, err := customerDao.Read()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(customers)
	})
}
