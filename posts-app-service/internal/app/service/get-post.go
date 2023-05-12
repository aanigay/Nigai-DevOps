package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"posts-app-service/internal/app/models"
)

func (s *Service) GetAll(c *gin.Context) {
	rows, err := s.db.Query("SELECT * FROM posts ORDER BY id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	posts := make([]models.Post, 0)
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UserId, pq.Array(&post.CommentsIds))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, post)
	}
	c.JSON(http.StatusOK, posts)
	err = rows.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (s *Service) GetById(c *gin.Context) {
	id := c.Query("id")
	row := s.db.QueryRow("SELECT * FROM posts WHERE id = $1", id)
	var post models.Post
	err := row.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UserId, pq.Array(&post.CommentsIds))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}
