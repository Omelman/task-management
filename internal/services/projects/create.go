package projects

import (
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo"
)

func CreateProject(newProject models.ProjectRequest) (models.Project, error) {
	project, err := repo.Get().Projects().Create(models.Project{
		ProjectName: newProject.Data.ProjectName,
		Description: newProject.Data.Description,
	})
	if err != nil {
		return models.Project{}, err
	}
	return project, nil
}
