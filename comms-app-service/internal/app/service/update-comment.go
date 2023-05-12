package service

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) UpdateComment(c *gin.Context) {
	r := struct {
		Id   int    `json:"id"`
		Body string `json:"body"`
	}{}
	if err := c.BindJSON(&r); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := s.db.Exec("UPDATE comments SET body = $1 WHERE id = $2", r.Body, r.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": r.Id})
}
