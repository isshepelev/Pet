package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func main() {
	_, err := NewPostgresDB()
	if err != nil {
		logrus.Fatalf("database initializing failed:%s", err.Error())
	}
	r := gin.Default()

	r.POST("/NewUser", func(c *gin.Context) {
		var input User
		if err := c.BindJSON(&input); err != nil {
			return
		}

		// Тело POST-запроса (в данном случае, JSON)
		// requestBody := []byte(`{"key": "value"}`)
		requestBody, err := json.Marshal(input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// URL Java-сервера
		javaURL := "http://localhost:8080/save"

		// Отправляем POST-запрос из Go на Java
		response, err := http.Post(javaURL, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("Error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer response.Body.Close()

		// Обработка ответа
		// Ваш код для обработки ответа от сервера Java

		c.JSON(http.StatusOK, gin.H{"message": "POST request sent to Java server"})
	})
	r.Run("127.0.0.1:80")
}

const (
	Host     = "localhost"
	Port     = "5432"
	Username = "postgres"
	Password = "postgres"
	DBName   = "pet"
)

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		Host, Port, Username, Password, DBName))
	if err != nil {
		return nil, err
	}

	return db, nil
}
