package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/hendrikTpl/go-psql-jwt-service/backend/models"
	"github.com/hendrikTpl/go-psql-jwt-service/backend/routes"
	"github.com/hendrikTpl/go-psql-jwt-service/backend/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Connect to PostgreSQL
	dbDSN := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", dbDSN)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	db.AutoMigrate(&models.User{}) // Migrate the User model

	// Connect to Redis
	redisAddr := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Error parsing REDIS_DB: %v", err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	// Get JWT secret from environment variable
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	// Create services
	userService := services.NewUserService(db)
	tokenService := services.NewTokenService(redisClient, jwtSecret)

	// Create router
	router := mux.NewRouter()

	// Set up API routes
	routes.SetUpRoutes(router, userService, tokenService)

	// Start the server
	fmt.Printf("Server listening on port 5001\n")
	log.Fatal(http.ListenAndServe(":5001", router))
}
