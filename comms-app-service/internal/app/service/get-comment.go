package service

import (
	"comms-app-service/internal/app/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

func (s *Service) GetAll(c *gin.Context) {
	rows, err := s.db.Query("SELECT * FROM comments")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	comments := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.Id, &comment.Body, &comment.CreatedAt, &comment.UserId, &comment.PostId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		comments = append(comments, comment)
	}
	c.JSON(http.StatusOK, comments)
	err = rows.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (s *Service) GetByPostId(c *gin.Context) {
	id := c.Query("postId")
	row := s.db.QueryRow("SELECT commentsids FROM posts WHERE id = $1", id)
	commentsIds := make([]int64, 0)
	err := row.Scan(pq.Array(&commentsIds))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	comments := make([]models.Comment, 0)
	for _, commentId := range commentsIds {
		row = s.db.QueryRow("SELECT * FROM comments WHERE id = $1", commentId)
		comment, err := models.RowToComment(row)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "debug": "get-comment.go::54"})
		}
		comments = append(comments, comment)
	}
	c.JSON(http.StatusOK, commentsIds)
}

func (s *Service) GetById(c *gin.Context) {
	id := c.Query("id")
	row := s.db.QueryRow("SELECT * FROM comments WHERE id = $1 ORDER BY id", id)
	var comment models.Comment
	err := row.Scan(&comment.Id, &comment.Body, &comment.CreatedAt, &comment.UserId, &comment.PostId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, comment)
}
