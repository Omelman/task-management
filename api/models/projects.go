package models

type Project struct {
	ID          int    `json:"id"`
	ProjectName string `json:"project_name"`
	Description string `json:"description"`
}
