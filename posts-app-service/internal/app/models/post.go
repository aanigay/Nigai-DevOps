package models

import "time"

type Post struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UserId      int64     `json:"user_id"`
	CommentsIds []int64   `json:"comments_ids"`
}
