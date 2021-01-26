package tasks

import "github.com/Omelman/task-management/internal/models"

type Tasks interface {
	Create(newTask models.Task) (models.Task, error)
	Delete(id int) error
	Update(t models.Task) (models.Task, error)
	GetAll(id int) ([]models.Task, error)
}
