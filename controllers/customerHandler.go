package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/felipehfs/testapi/models"
	"github.com/gorilla/mux"
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

		if err := customer.IsValid(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdCustomer, err := customerDao.Create(customer)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdCustomer)
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

// UpdateCustomer changes the customer by id
func UpdateCustomer(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customerDao := models.NewCustomerDao(db)
		vars := mux.Vars(r)
		var customer models.Customer

		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := customerDao.Update(vars["id"], customer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(customer)
	})
}

// RemoveCustomer removes the customer in the database
func RemoveCustomer(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customerDao := models.NewCustomerDao(db)
		vars := mux.Vars(r)
		if err := customerDao.Remove(vars["id"]); err != nil {
			var statusCode int

			if err == sql.ErrNoRows {
				statusCode = http.StatusNotFound
			} else {
				statusCode = http.StatusInternalServerError
			}

			http.Error(w, err.Error(), statusCode)
		}
	})
}
