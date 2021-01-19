package context

import (
	"context"
	"github.com/Omelman/task-management/api/repo/postgres/projects"
	"net/http"
)

type ctxKey int

const (
	projectsCtxKey ctxKey = iota
)

func CtxProject(entry projects.Projects) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, projectsCtxKey, entry)
	}
}

func Project(r *http.Request) projects.Projects {
	return r.Context().Value(projectsCtxKey).(projects.Projects)
}
