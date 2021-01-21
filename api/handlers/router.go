package handlers

import (
	handler "github.com/Omelman/task-management/api/handlers/projects"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.CreateProject)
	})

	return r
}
