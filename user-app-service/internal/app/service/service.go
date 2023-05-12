package service

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
)

type Service struct {
	engine *gin.Engine
	dbUrl  string
	db     *sql.DB
}

func NewService() *Service {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		fmt.Println("NULL ENV DBSTRING")
		dbUrl = "postgres://root:root@localhost:5432/app-db?sslmode=disable"
	}

	return &Service{
		engine: gin.Default(),
		dbUrl:  dbUrl,
	}
}

func (s *Service) Run() error {

	db, err := sql.Open("postgres",
		s.dbUrl)
	if err != nil {
		return err
	}
	s.db = db
	if err = s.db.Ping(); err != nil {
		return err
	}

	s.engine.Use(CORS())
	s.engine.GET("/get-all", s.getAll)
	s.engine.GET("/getById", s.getById)
	s.engine.POST("/login", s.login)
	s.engine.POST("/register", s.register)

	err = s.engine.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return err
	}
	return nil
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
