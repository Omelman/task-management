package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("OK"))
		})
	})

	return r
}
