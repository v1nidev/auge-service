package main

import (
	"log"
	"fmt"
	"os"
	"context"

	"github.com/v1nidev/auge-service/ent"
	"github.com/v1nidev/auge-service/ent/migrate"

	_ "github.com/lib/pq"
)

const defaultPort = "3000"

func main() {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbConnectionString := fmt.Sprintf("host=db port=5432 user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword)

	client, err := ent.Open("postgres", dbConnectionString)
	if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
	}
}
