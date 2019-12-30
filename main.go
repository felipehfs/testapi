package main

import (
	"log"
	"net/http"

	"github.com/felipehfs/testapi/config"
	"github.com/felipehfs/testapi/routes"
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

	log.Fatal(http.ListenAndServe(":8080", r))
}
