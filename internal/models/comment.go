package models

type CommentRequest struct {
	Data Comment `json:"data"`
}

type CommentListRequest struct {
	Data []Comment `json:"data"`
}

type Comment struct {
	ID          int    `json:"id"`
	CommentText string `json:"comment_text"`
	TaskID      int    `json:"task_id"`
}
