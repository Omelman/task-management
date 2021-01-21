package repo

import "github.com/Omelman/task-management/api/repo/postgres/projects"

type Repo interface {
	Projects() projects.Projects
}
