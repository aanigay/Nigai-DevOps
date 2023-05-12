package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	Id        int64     `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UserId    int64     `json:"user_id"`
	PostId    int64     `json:"post_id"`
}

func RowToComment(r *sql.Row) (Comment, error) {
	var comment Comment
	err := r.Scan(&comment.Id, &comment.Body, &comment.CreatedAt, &comment.UserId, &comment.PostId)
	return comment, err
}
