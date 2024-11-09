package models

import "database/sql"

type Todo struct {
	TodoId    int          `json:"todo_id"`
	Content   string       `json:"content"`
	IsDone    bool         `json:"is_done"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type ContentData struct {
	Content string `json:"content" validate:"required"`
}
