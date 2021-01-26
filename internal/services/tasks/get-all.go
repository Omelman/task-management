package tasks

import (
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo"
)

func GetTasks(columnId int) ([]models.Task, error) {
	tasks, err := repo.Get().Tasks().GetAll(columnId)
	if err != nil {
		return []models.Task{}, err
	}
	return tasks, nil
}
