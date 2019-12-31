package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/felipehfs/testapi/models"
)

// CreateOrder inserts the new orders with processing status
func CreateOrder(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orderDao := models.NewOrderDao(db)
		var request models.Order

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user := r.Context().Value("user")
		claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
		activeID := int64(claims["id"].(float64))
		request.Author = activeID

		created, err := orderDao.Create(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(created)
	})
}

// ReadOrder retrieves all orders
func ReadOrder(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orderDao := models.NewOrderDao(db)

		user := r.Context().Value("user")
		claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
		id := int64(claims["id"].(float64))

		orders, err := orderDao.Read(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(orders)
	})
}
