package main

import (
	"log"
	"net/http"

	"github.com/felipehfs/testapi/config"
	"github.com/felipehfs/testapi/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	conn, err := config.CreateConnection()

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		Debug:            true,
		AllowedHeaders:   []string{"Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	routes.Routes(conn, r)

	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}
