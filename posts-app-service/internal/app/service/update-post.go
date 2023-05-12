package service

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) UpdatePost(c *gin.Context) {
	r := struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}{}
	if err := c.BindJSON(&r); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := s.db.Exec("UPDATE posts SET title = $1, body = $2 WHERE id = $3", r.Title, r.Body, r.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": r.Id})
}
