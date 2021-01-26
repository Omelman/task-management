package comments

import (
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo"
)

func GetComments(taskId int) ([]models.Comment, error) {
	comments, err := repo.Get().Comments().GetAll(taskId)
	if err != nil {
		return []models.Comment{}, err
	}
	return comments, nil
}
