package repo

import (
	"github.com/Omelman/task-management/internal/repo/postgres/columns"
	"github.com/Omelman/task-management/internal/repo/postgres/projects"
	"sync"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	projects projects.Projects
	columns  columns.Columns
}

func Get() Repo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			projects: projects.NewProjects(),
			columns:  columns.NewColumns(),
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
