package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type server struct {
	router *chi.Mux
	db     *badger.DB
}

func (s *server) middlewares() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
}

func (s *server) routes() {
	s.router.Get("/v1/ping", s.handlePing())
}

func (s *server) handlePing() http.HandlerFunc {
	type response struct {
		Pong bool `json:"pong"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(response{Pong: true})
	}
}

func newDB(path string) *badger.DB {
	db, err := badger.Open(badger.DefaultOptions(path))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db
}

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	srv := server{
		router: chi.NewRouter(),
		db:     newDB("/tmp/badger"),
	}

	srv.middlewares()
	srv.routes()

	log.Fatal(http.ListenAndServe(host+":"+port, srv.router))
}
