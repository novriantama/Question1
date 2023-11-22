package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/novriantama/question1/pkg/handlers"
	"github.com/novriantama/question1/pkg/repository"
	"github.com/novriantama/question1/pkg/services"
	"github.com/novriantama/question1/pkg/sqlc/db"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type UserPayload struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	dbURL := os.Getenv("DATABASE_URL")

	// Database connection
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		// Handle database connection error
		panic(err)
	}
	defer conn.Close(context.Background())
	queries := db.New(conn)

	// Repository
	repo := repository.NewRepository(conn, queries)

	// Service
	service := services.NewService(repo)

	// Handlers
	handler := handlers.NewHandlers(service)

	// Gin setup
	r := gin.Default()

	// Routes
	r.POST("/api/users", handler.CreateUser)
	r.POST("/api/users/generateotp", handler.GenerateOtp)
	r.POST("/api/users/verifyotp", handler.VerifyOtp)
	r.GET("/users/:id", handler.GetUserByID)

	// Run the server
	r.Run("localhost:8080")
}
