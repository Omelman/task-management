package projects

import (
	"database/sql"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo/postgres"
	"github.com/pkg/errors"
)

type projects struct {
	db *sql.DB
}

func NewProjects() *projects {
	return &projects{
		db: postgres.GetDB().DB(),
	}
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

func (q *projects) Update(newProject models.Project) (models.Project, error) {
	query := "UPDATE projects SET project_name = $1, description = $2 WHERE id = $3;"
	res, err := q.db.Exec(query, newProject.ProjectName, newProject.Description, newProject.ID)

	if err != nil {
		return models.Project{}, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return models.Project{}, err
	} else if rowsCount == 0 {
		return models.Project{}, errors.New("not found")
	}

	return newProject, nil
}

func (q *projects) Delete(id int) error {
	res, err := q.db.Exec("DELETE FROM projects WHERE id = $1;", id)

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

func (q *projects) GetAll() ([]models.Project, error) {
	rows, err := q.db.Query("SELECT * FROM projects;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projectList []models.Project
	project := models.Project{}
	for rows.Next() {
		if err = rows.Scan(&project.ID, &project.ProjectName, &project.Description); err != nil {
			return nil, err
		}
		projectList = append(projectList, project)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projectList, nil
}
