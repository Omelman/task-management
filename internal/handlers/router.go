package handlers

import (
	"github.com/Omelman/task-management/internal/handlers/columns"
	"github.com/Omelman/task-management/internal/handlers/comments"
	"github.com/Omelman/task-management/internal/handlers/projects"
	"github.com/Omelman/task-management/internal/handlers/tasks"
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
			r.Post("/", tasks.CreateTask)
			r.Put("/", tasks.UpdateTask)
			r.Delete("/{id}", tasks.DeleteTask)
			r.Get("/{column_id}", tasks.GetTasks)
		})

		r.Route("/comment", func(r chi.Router) {
			r.Post("/", comments.CreateComment)
			r.Put("/", comments.UpdateComment)
			r.Delete("/{id}", comments.DeleteComment)
			r.Get("/{task_id}", comments.GetComments)
		})
	})

	return r
}
