package repo

import (
	"github.com/Omelman/task-management/internal/repo/postgres/columns"
	"github.com/Omelman/task-management/internal/repo/postgres/comments"
	"github.com/Omelman/task-management/internal/repo/postgres/projects"
	"github.com/Omelman/task-management/internal/repo/postgres/tasks"
)

type Repo interface {
	Projects() projects.Projects
	Columns() columns.Columns
	Tasks() tasks.Tasks
	Comments() comments.Comments
}
