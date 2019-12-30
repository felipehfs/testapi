package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "123456"
	DB_NAME     = "testapi"
)

// CreateConnection creates the connection
func CreateConnection() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	return sql.Open("postgres", dbInfo)
}
