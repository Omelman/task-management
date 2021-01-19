package models

type ProjectRequest struct {
	ID          int    `json:"id"`
	ProjectName string `json:"project_name"`
	Description string `json:"description"`
}
