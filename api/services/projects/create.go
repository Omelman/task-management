package projects

import (
	"github.com/Omelman/task-management/api/models"
	"github.com/Omelman/task-management/api/repo"
)

func CreateProject(newProject models.ProjectRequest) (models.Project, error) {
	project, err := repo.Get().Projects().Create(models.Project{
		ID:          newProject.ID,
		ProjectName: newProject.ProjectName,
		Description: newProject.Description,
	})
	if err != nil {
		return models.Project{}, err
	}
	return project, nil
}
