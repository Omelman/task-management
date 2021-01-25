package handlers

import (
	"github.com/Omelman/task-management/internal/handlers/columns"
	"github.com/Omelman/task-management/internal/handlers/projects"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		r.Route("/project", func(r chi.Router) {
			r.Post("/", projects.CreateProject)
			r.Put("/", projects.UpdateProject)
			r.Get("/", projects.GetProjects)
			r.Delete("/{id}", projects.DeleteProject)
		})

		r.Route("/column", func(r chi.Router) {
			r.Post("/", columns.CreateColumn)
			r.Put("/", columns.UpdateColumn)
			r.Get("/{project_id}", columns.GetColumns)
			r.Delete("/{id}", columns.DeleteColumn)
		})

		r.Route("/task", func(r chi.Router) {

		})

		r.Route("/comment", func(r chi.Router) {

		})
	})

	return r
}
