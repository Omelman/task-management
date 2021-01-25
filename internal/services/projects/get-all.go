package projects

import (
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo"
)

func GetProjects() ([]models.Project, error) {
	projects, err := repo.Get().Projects().GetAll()
	if err != nil {
		return []models.Project{}, err
	}
	return projects, nil
}
