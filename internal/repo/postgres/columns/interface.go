package columns

import "github.com/Omelman/task-management/internal/models"

type Columns interface {
	Create(newColumn models.Column) (models.Column, error)
	Delete(id int) error
	Update(newColumn models.Column) (models.Column, error)
	GetAll(projectID int) ([]models.Column, error)
}
