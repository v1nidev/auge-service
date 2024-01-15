package database

import (
	"fmt"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

//go:generate go run gen.go
func Open() (*sql.DB, error) {
	dbConnectionString := fmt.Sprintf(
		"host=db port=5432 user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"),
	)

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	return db, nil
}
