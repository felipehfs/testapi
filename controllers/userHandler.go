package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"

	"github.com/dgrijalva/jwt-go"
	"github.com/felipehfs/testapi/models"
)

const (
	APP_KEY = "s3cr3t"
)

func AuthMiddleware(next http.Handler) http.Handler {
	if len(APP_KEY) == 0 {
		log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
	}

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(APP_KEY), nil
		},
	})

	return jwtMiddleware.Handler(next)
}

// Register inserts the new user
func Register(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		userDao := models.NewUserDao(db)
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Erro na sintaxe do JSON", http.StatusBadRequest)
			return
		}

		if err := userDao.Create(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":  user.Name,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			"iat":   time.Now().Unix(),
		})

		tokenString, err := token.SignedString([]byte(APP_KEY))
		if err != nil {
			http.Error(w, "Generated token failed", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"token": tokenString,
			"email": user.Email,
		}

		json.NewEncoder(w).Encode(response)
	})
}

func Login(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request map[string]string
		userDao := models.NewUserDao(db)

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := userDao.Find(request["email"], request["password"])

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			"iat":   time.Now().Unix(),
		})

		tokenString, err := token.SignedString([]byte(APP_KEY))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"token": tokenString,
			"email": user.Email,
		}

		json.NewEncoder(w).Encode(response)
	})
}
