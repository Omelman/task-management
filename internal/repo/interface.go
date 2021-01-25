package repo

import (
	"github.com/Omelman/task-management/internal/repo/postgres/columns"
	"github.com/Omelman/task-management/internal/repo/postgres/projects"
)

type Repo interface {
	Projects() projects.Projects
	Columns() columns.Columns
}
