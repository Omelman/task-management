package repo

import (
	"github.com/Omelman/task-management/api/repo/postgres/projects"
	"sync"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	projects projects.Projects
}

func Get() Repo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			projects: projects.NewProjects(),
		}
	})
	return err
}

func (r postgresRepo) Projects() projects.Projects {
	return r.projects
}
