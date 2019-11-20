package util

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// GetDB method returns a DB instance
func GetDB() (*sql.DB, error) {
	connectionString := "user=postgres password=password dbname=QuikNote"
	// connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connectionString == "" {
		return nil, errors.New("'POSTGRES_CONNECTION_STRING' environment variable not set")
	}
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}
	return conn, nil
}
