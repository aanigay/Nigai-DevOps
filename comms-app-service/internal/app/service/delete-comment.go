package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) DeleteById(c *gin.Context) {
	id := c.Query("id")
	_, err := s.db.Exec("DELETE FROM comments WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = s.db.Exec("UPDATE posts SET commentsids = array_remove(commentsids, $1)", id)
	c.JSON(http.StatusOK, gin.H{"id": id})
}
