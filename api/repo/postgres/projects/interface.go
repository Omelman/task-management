package projects

import (
	"github.com/Omelman/task-management/api/models"
)

type Projects interface {
	Create(newProject models.Project) (models.Project, error)
}
