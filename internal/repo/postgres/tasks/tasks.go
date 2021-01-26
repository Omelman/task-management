package tasks

import (
	"database/sql"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo/postgres"
	"github.com/pkg/errors"
)

type tasks struct {
	db *sql.DB
}

func NewTasks() *tasks {
	return &tasks{
		db: postgres.GetDB().DB(),
	}
}

func (q *tasks) Create(newTask models.Task) (models.Task, error) {
	query := "INSERT INTO tasks (name, description, index, column_id) VALUES ($1, $2, $3, $4) RETURNING id;"
	row := q.db.QueryRow(query, newTask.Name, newTask.Description, newTask.Index, newTask.ColumnID)

	var id int
	if err := row.Scan(&id); err != nil {
		return models.Task{}, err
	}
	newTask.ID = id

	return newTask, nil
}

func (q *tasks) Delete(id int) error {
	res, err := q.db.Exec("DELETE FROM tasks WHERE id = $1;", id)

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

func (q *tasks) Update(task models.Task) (models.Task, error) {
	query := "UPDATE tasks SET name = $1, description = $2, index = $3, column_id = $4 WHERE id = $5;"
	res, err := q.db.Exec(query, task.Name, task.Description, task.Index, task.ColumnID, task.ID)

	if err != nil {
		return models.Task{}, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return models.Task{}, err
	} else if rowsCount == 0 {
		return models.Task{}, errors.New("not found")
	}

	return task, nil
}

func (q *tasks) GetAll(id int) ([]models.Task, error) {
	rows, err := q.db.Query("SELECT tasks.* FROM tasks INNER JOIN columns ON "+
		"tasks.column_id = columns.id WHERE column_id = $1;", id)
	if err != nil {
		return nil, err
	} else if !rows.Next() {
		return nil, errors.New("not found")
	}

	var ts []models.Task
	t := models.Task{}
	for rows.Next() {
		if err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.Index, &t.ColumnID); err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ts, nil
}
