package routes

import (
	"database/sql"

	"github.com/felipehfs/testapi/controllers"
	"github.com/gorilla/mux"
)

// Routes setup the routes to the api
func Routes(conn *sql.DB, r *mux.Router) {

	findProduct := controllers.AuthMiddleware(controllers.FindProduct(conn))
	readProduct := controllers.AuthMiddleware(controllers.ReadProduct(conn))
	updateProduct := controllers.AuthMiddleware(controllers.UpdateProduct(conn))
	removeProduct := controllers.AuthMiddleware(controllers.RemoveProduct(conn))
	createProduct := controllers.AuthMiddleware(controllers.CreateProduct(conn))

	createCustomer := controllers.AuthMiddleware(controllers.CreateCustomer(conn))
	readCustomer := controllers.AuthMiddleware(controllers.ReadCustomer(conn))
	updateCustomer := controllers.AuthMiddleware(controllers.UpdateCustomer(conn))
	removeCustomer := controllers.AuthMiddleware(controllers.RemoveCustomer(conn))

	createOrder := controllers.AuthMiddleware(controllers.CreateOrder(conn))
	readOrder := controllers.AuthMiddleware(controllers.ReadOrder(conn))
	updateOrder := controllers.AuthMiddleware(controllers.UpdateOrder(conn))
	findOrder := controllers.AuthMiddleware(controllers.FindOrder(conn))

	r.Handle("/api/login", controllers.Login(conn)).Methods("POST")
	r.Handle("/api/register", controllers.Register(conn)).Methods("POST")

	r.Handle("/api/products", createProduct).Methods("POST")
	r.Handle("/api/products", readProduct).Methods("GET")
	r.Handle("/api/products/{id}", updateProduct).Methods("PUT")
	r.Handle("/api/products/{id}", removeProduct).Methods("DELETE")
	r.Handle("/api/products/{id}", findProduct).Methods("GET")

	r.Handle("/api/customers", createCustomer).Methods("POST")
	r.Handle("/api/customers", readCustomer).Methods("GET")
	r.Handle("/api/customers/{id}", updateCustomer).Methods("PUT")
	r.Handle("/api/customers/{id}", removeCustomer).Methods("DELETE")

	r.Handle("/api/orders", createOrder).Methods("POST")
	r.Handle("/api/orders", readOrder).Methods("GET")
	r.Handle("/api/orders/{id}", updateOrder).Methods("PUT")
	r.Handle("/api/orders/{id}", findOrder).Methods("GET")
}
