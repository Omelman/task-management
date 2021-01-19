package projects

import (
	"database/sql"
	"github.com/Omelman/task-management/api/models"
)

type projects struct {
	db *sql.DB
}

func NewProjects(db *sql.DB) *projects {
	return &projects{db: db}
}

func (q *projects) Create(newProject models.Project) (models.Project, error) {
	query := "INSERT INTO projects (project_name, description) VALUES ($1, $2) RETURNING id;"
	row := q.db.QueryRow(query, newProject.ProjectName, newProject.Description)

	var id int
	if err := row.Scan(&id); err != nil {
		return models.Project{}, err
	}
	newProject.ID = id

	return newProject, nil
}
