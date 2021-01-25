package projects

import (
	"github.com/Omelman/task-management/internal/models"
)

type Projects interface {
	Create(newProject models.Project) (models.Project, error)
	Update(newProject models.Project) (models.Project, error)
	Delete(id int) error
	GetAll() ([]models.Project, error)
}
