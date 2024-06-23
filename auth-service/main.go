package main

import (
	applications "authService/applications/services"
	handler "authService/handlers"
	infrastructure "authService/infrastructure/dbMigrations"
	"authService/infrastructure/messaging"
	postgres "authService/infrastructure/repository/postgres"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	// Initialize database
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate user table
	infrastructure.AutoMigrate(db)

	// Initialize NATS
	nc := messaging.NewNATSClient()

	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)

	// Initialize services
	userService := applications.NewUserService(userRepo, nc)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)

	// Setup routes
	app.Post("/register", userHandler.Register)

	// Start the application
	log.Fatal(app.Listen(":3001"))
}
