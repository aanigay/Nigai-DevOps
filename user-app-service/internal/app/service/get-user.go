package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-app-service/internal/app/models"
)

func (s *Service) getAll(c *gin.Context) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
	err = rows.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (s *Service) getById(c *gin.Context) {
	userId := c.Query("id")
	row := s.db.QueryRow("SELECT id, name, role FROM users WHERE id = $1", userId)
	var user struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Role string `json:"role"`
	}
	err := row.Scan(&user.Id, &user.Name, &user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
