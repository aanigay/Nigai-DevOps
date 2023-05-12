package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) deleteById(c *gin.Context) {
	id := c.Query("id")
	s.db.QueryRow("DELETE FROM posts WHERE id = $1", id)
	c.JSON(http.StatusOK, gin.H{"id": id})
}
