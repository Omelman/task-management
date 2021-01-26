package comments

import (
	"database/sql"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo/postgres"
	"github.com/pkg/errors"
)

type comments struct {
	db *sql.DB
}

func NewComments() *comments {
	return &comments{
		db: postgres.GetDB().DB(),
	}
}

func (q *comments) Create(newComment models.Comment) (models.Comment, error) {
	query := "INSERT INTO comments (comment_text, task_id) VALUES ($1, $2) RETURNING id;"
	row := q.db.QueryRow(query, newComment.CommentText, newComment.TaskID)

	var id int
	if err := row.Scan(&id); err != nil {
		return models.Comment{}, err
	}
	newComment.ID = id

	return newComment, nil
}

func (q *comments) Delete(id int) error {
	res, err := q.db.Exec("DELETE FROM comments WHERE id = $1;", id)

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

func (q *comments) Update(comment models.Comment) (models.Comment, error) {
	query := "UPDATE comments SET text = $1, task_id = $2 WHERE id = $3;"
	res, err := q.db.Exec(query, comment.CommentText, comment.TaskID, comment.ID)

	if err != nil {
		return models.Comment{}, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return models.Comment{}, err
	} else if rowsCount == 0 {
		return models.Comment{}, errors.New("not found")
	}

	return comment, err
}

func (q *comments) GetAll(id int) ([]models.Comment, error) {
	rows, err := q.db.Query("SELECT * FROM tasks WHERE id = $1;", id)
	if err != nil {
		return nil, err
	} else if !rows.Next() {
		return nil, errors.New("not found")
	}

	rows, err = q.db.Query("SELECT * FROM comments WHERE task_id = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	comment := models.Comment{}
	for rows.Next() {
		if err := rows.Scan(&comment.ID, &comment.CommentText, &comment.TaskID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
