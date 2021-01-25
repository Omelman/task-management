package models

type CommentRequest struct {
	Data Comment `json:"data"`
}

type CommentListRequest struct {
	Data []Comment `json:"data"`
}

type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	TaskID int    `json:"task_id"`
}
