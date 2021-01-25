package columns

import (
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo"
)

func GetColumns(projectId int) ([]models.Column, error) {
	columns, err := repo.Get().Columns().GetAll(projectId)
	if err != nil {
		return []models.Column{}, err
	}
	return columns, nil
}
