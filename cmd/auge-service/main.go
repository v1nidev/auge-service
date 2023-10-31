package main

import (
	"log"
	"os"
	"net/http"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/v1nidev/auge-service/database"
	"github.com/v1nidev/auge-service/internal/member"
	"github.com/v1nidev/auge-service/internal/package"
)

const defaultPort = "3000"

func Routes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
	)

	router.Mount("/member", member.Routes())
	router.Mount("/package", gymPackage.Routes(gymPackage.ConfigPackageDi(db)))

	return router
}

func main() {
	db, err := database.Open()
	defer db.Close()
	if err != nil {
		os.Exit(1)
	}

	router := Routes(db)
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":3000", router))


}

