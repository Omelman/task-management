package models

type TaskRequest struct {
	Data Task `json:"data"`
}

type TaskListRequest struct {
	Data []Task `json:"data"`
}

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Index       int    `json:"index"`
	ColumnID    int    `json:"column_id"`
}
