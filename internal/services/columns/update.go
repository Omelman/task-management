package columns

import (
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo"
)

func UpdateColumn(newColumn models.ColumnRequest) (models.Column, error) {
	column, err := repo.Get().Columns().Update(models.Column{
		ID:         newColumn.Data.ID,
		ColumnName: newColumn.Data.ColumnName,
		Index:      newColumn.Data.Index,
		ProjectID:  newColumn.Data.ProjectID,
	})
	if err != nil {
		return models.Column{}, err
	}
	return column, nil
}
