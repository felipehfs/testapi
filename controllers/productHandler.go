package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/felipehfs/testapi/models"
	"github.com/gorilla/mux"
)

// CreateProduct handler
func CreateProduct(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var product models.Product
		productDao := models.NewProductDao(db)

		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			log.Println(err)
			http.Error(w, "erro na sintaxe do JSON", http.StatusBadRequest)
			return
		}

		createdProduct, err := productDao.Create(product)

		if err != nil {
			log.Println(err)
			http.Error(w, "erro na criação do produto", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(createdProduct)
	})
}

// ReadProduct handler
func ReadProduct(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		productDao := models.NewProductDao(db)
		products, err := productDao.Read()
		if err != nil {
			log.Println(err)
			http.Error(w, "um erro ocorreu ao tentar recuperar os produtos", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(products)
	})
}

// UpdateProduct change the product
func UpdateProduct(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		productDao := models.NewProductDao(db)
		var p models.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "erro em interpretar o json", http.StatusBadRequest)
			return
		}

		if err := productDao.Update(p, vars["id"]); err != nil {
			http.Error(w, "erro em atualizar o produto", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

// RemoveProduct removes the product in the database
func RemoveProduct(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		productDao := models.NewProductDao(db)
		id, err := strconv.ParseInt(vars["id"], 10, 64)

		if err != nil {
			http.Error(w, "Parâmetro inválido", http.StatusBadRequest)
			return
		}

		if err := productDao.Remove(id); err != nil {
			log.Println(err)
			http.Error(w, "um erro ocorreu ao remover o produto", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

// FindProduct searches the product by ID
func FindProduct(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		productDao := models.NewProductDao(db)
		id, err := strconv.ParseInt(vars["id"], 10, 64)

		if err != nil {
			http.Error(w, "Parâmetro inválido", http.StatusBadRequest)
			return
		}
		product, err := productDao.Find(id)

		if err != nil {
			log.Println(err)
			switch err {
			case sql.ErrNoRows:
				http.Error(w, "Id não encontrado", http.StatusNotFound)
				return
			default:
				http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
				return
			}
		}

		json.NewEncoder(w).Encode(product)
	})
}
