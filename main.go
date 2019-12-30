package main

import (
	"log"
	"net/http"

	"github.com/felipehfs/testapi/config"
	"github.com/felipehfs/testapi/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	conn, err := config.CreateConnection()

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	r := mux.NewRouter()
	routes.Routes(conn, r)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Request-With"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))
}
