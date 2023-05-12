package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) changeRole(c *gin.Context) {
	r := struct {
		Id   int    `json:"id"`
		Role string `json:"role"`
	}{}
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := s.db.Exec("UPDATE users SET role = $1 WHERE id = $2", r.Role, r.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
