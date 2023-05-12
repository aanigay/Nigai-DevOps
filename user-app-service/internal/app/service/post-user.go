package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"time"
	"user-app-service/internal/app/models"
)

const (
	secretKey = "secret"
)

type MyJWTClaims struct {
	Id   string `json:"id"`
	Name string `json:"Name"`
	jwt.RegisteredClaims
}

func (s *Service) login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	row := s.db.QueryRow("SELECT * FROM users WHERE name = $1 AND password = $2", user.Name, user.Password)
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Role)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No such user!"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		Id:   strconv.Itoa(int(user.Id)),
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.Id)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": ss, "user_id": user.Id, "name": user.Name, "role": user.Role})
}

func (s *Service) register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.db.QueryRow("INSERT INTO users (name, password, role) VALUES ($1, $2, $3) RETURNING id", &user.Name, &user.Password, "editor").Scan(&user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists!"})
		return
	}
	user.Role = "editor"
	c.JSON(http.StatusOK, user)
}
