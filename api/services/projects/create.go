package projects

import (
	"github.com/Omelman/task-management/api/context"
	"github.com/Omelman/task-management/api/models"
	"net/http"
)

func CreateProject(r *http.Request, newProject models.ProjectRequest) (models.Project, error) {
	project, err := context.Project(r).Create(models.Project{
		ID:          newProject.ID,
		ProjectName: newProject.ProjectName,
		Description: newProject.Description,
	})
	if err != nil {
		return models.Project{}, err
	}

	return project, nil
}
