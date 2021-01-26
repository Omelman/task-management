package comments

import "github.com/Omelman/task-management/internal/models"

type Comments interface {
	Create(newComment models.Comment) (models.Comment, error)
	Delete(id int) error
	Update(comment models.Comment) (models.Comment, error)
	GetAll(id int) ([]models.Comment, error)
}
