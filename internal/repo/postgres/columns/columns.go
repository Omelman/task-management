package columns

import (
	"database/sql"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo/postgres"
	"github.com/pkg/errors"
)

type columns struct {
	db *sql.DB
}

func NewColumns() *columns {
	return &columns{
		db: postgres.GetDB().DB(),
	}
}

func (q *columns) Create(newColumn models.Column) (models.Column, error) {
	query := "INSERT INTO columns (column_name, index, project_id) VALUES ($1, $2, $3) RETURNING id;"
	row := q.db.QueryRow(query, newColumn.ColumnName, newColumn.Index, newColumn.ProjectID)

	var id int
	if err := row.Scan(&id); err != nil {
		return models.Column{}, err
	}
	newColumn.ID = id

	return newColumn, nil
}

func (q *columns) Delete(id int) error {
	res, err := q.db.Exec("DELETE FROM columns WHERE id = $1;", id)

	if err != nil {
		return err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsCount == 0 {
		return errors.New("not found")
	}

	return nil
}

func (q *columns) Update(newColumn models.Column) (models.Column, error) {
	query := "UPDATE columns SET column_name = $1, index = $2, project_id = $3 WHERE id = $4;"
	res, err := q.db.Exec(query, newColumn.ColumnName, newColumn.Index, newColumn.ProjectID, newColumn.ID)

	if err != nil {
		return models.Column{}, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return models.Column{}, err
	} else if rowsCount == 0 {
		return models.Column{}, errors.New("not found")
	}

	return newColumn, nil
}

func (q *columns) GetAll(projectID int) ([]models.Column, error) {
	rows, err := q.db.Query("SELECT * FROM columns WHERE project_id = $1;", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []models.Column
	column := models.Column{}
	for rows.Next() {
		if err = rows.Scan(&column.ID, &column.ColumnName, &column.Index, &column.ProjectID); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}
