package database

import (
	"os"
	"fmt"

	"github.com/go-jet/jet/generator/postgres"
)

func Generate() (err error) {
	err = postgres.Generate("./database/gen", // or GenerateDSN(...)
		postgres.DBConnection{
			Host:       "db",
			Port:       5432,
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			SchemaName: "public",
			SslMode:    "disable",
	})

	if err != nil {
		return fmt.Errorf("could not generate database files: %w", err)
	}

	return err
}
