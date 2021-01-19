package handlers

import (
	"github.com/Omelman/task-management/api/context"
	"github.com/Omelman/task-management/api/handlers/middlewares"
	handler "github.com/Omelman/task-management/api/handlers/projects"
	"github.com/Omelman/task-management/api/repo/postgres"
	"github.com/Omelman/task-management/api/repo/postgres/projects"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(db postgres.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middlewares.CtxMiddleWare(
		context.CtxProject(projects.NewProjects(db.RawDB())),
	))

	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.CreateProject)
	})

	return r
}
