package main

import (
	"log"
	"fmt"
	"os"
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/v1nidev/auge-service/ent"
	"github.com/v1nidev/auge-service/ent/migrate"
	"github.com/v1nidev/auge-service/internal/member"

	_ "github.com/lib/pq"
)

const defaultPort = "3000"

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
	)

	router.Mount("/member", member.Routes())

	return router
}

func main() {
	router := Routes()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":3000", router))

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

