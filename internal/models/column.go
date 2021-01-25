package models

type ColumnRequest struct {
	Data Column `json:"data"`
}

type ColumnListRequest struct {
	Data []Column `json:"data"`
}

type Column struct {
	ID         int    `json:"id,omitempty"`
	ColumnName string `json:"column_name"`
	Index      int    `json:"index"`
	ProjectID  int    `json:"project_id"`
}
