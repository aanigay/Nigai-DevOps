package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) getRoleById(c *gin.Context) {
	id := c.Query("id")
	row := s.db.QueryRow("SELECT role FROM users WHERE id = $1", id)
	var role string
	err := row.Scan(&role)
	if err != nil {
		_, err := s.db.Exec("INSERT INTO users (id, role) VALUES ($1, $2)", id, "editor")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		role = "editor"
		c.JSON(http.StatusOK, gin.H{"role": role})
		return
	}
	c.JSON(http.StatusOK, gin.H{"role": role})
}
