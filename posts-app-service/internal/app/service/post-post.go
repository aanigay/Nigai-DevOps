package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"time"
)

func (s *Service) CreatePost(c *gin.Context) {
	var post struct {
		UserId int    `json:"user_id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}
	if err := c.BindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	CreatedAt := time.Now()
	var id int64
	err := s.db.QueryRow("INSERT INTO posts (title, body, createdAt, userId, commentsIds) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		post.Title, post.Body, CreatedAt, post.UserId, pq.Array(make([]string, 0))).Scan(&id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}
