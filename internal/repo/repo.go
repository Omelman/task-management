package repo

import (
	"github.com/Omelman/task-management/internal/repo/postgres/columns"
	"github.com/Omelman/task-management/internal/repo/postgres/comments"
	"github.com/Omelman/task-management/internal/repo/postgres/projects"
	"github.com/Omelman/task-management/internal/repo/postgres/tasks"
	"sync"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	projects projects.Projects
	columns  columns.Columns
	tasks    tasks.Tasks
	comments comments.Comments
}

func Get() Repo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			projects: projects.NewProjects(),
			columns:  columns.NewColumns(),
			tasks:    tasks.NewTasks(),
			comments: comments.NewComments(),
		}
	})
	return err
}

func (r postgresRepo) Projects() projects.Projects {
	return r.projects
}

func (r postgresRepo) Columns() columns.Columns {
	return r.columns
}

func (r postgresRepo) Tasks() tasks.Tasks {
	return r.tasks
}

func (r postgresRepo) Comments() comments.Comments {
	return r.comments
}
