package tasks

import (
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo"
)

func UpdateTask(newTask models.TaskRequest) (models.Task, error) {
	task, err := repo.Get().Tasks().Update(models.Task{
		ID:          newTask.Data.ID,
		Name:        newTask.Data.Name,
		Description: newTask.Data.Description,
		Index:       newTask.Data.Index,
		ColumnID:    newTask.Data.ColumnID,
	})
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
