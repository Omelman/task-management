package models

type ProjectRequest struct {
	Data Project `json:"data"`
}

type ProjectListRequest struct {
	Data []Project `json:"data"`
}

type Project struct {
	ID          int    `json:"id,omitempty"`
	ProjectName string `json:"project_name"`
	Description string `json:"description"`
}
