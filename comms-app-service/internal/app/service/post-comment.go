package service

import (
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Service) CreateComment(c *gin.Context) {
	var comment struct {
		Body   string `json:"body"`
		UserId int64  `json:"user_id"`
		PostId int64  `json:"post_id"`
	}
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	CreatedAt := time.Now()
	var id int64
	err := s.db.QueryRow("INSERT INTO comments (body, createdAt, userId, postId) VALUES ($1, $2, $3, $4) RETURNING id",
		comment.Body, CreatedAt, comment.UserId, comment.PostId).Scan(&id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = s.db.Exec("UPDATE posts SET commentsids = array_append(commentsids, $1) WHERE id = $2", id, comment.PostId)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"id": id})
}
