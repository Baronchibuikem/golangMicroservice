package main

import (
	"log"
	"mailingService/applications/services"
	rest "mailingService/handlers"
	"mailingService/infrastructure/messaging"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize NATS
	nc := messaging.NewNATSClient()

	// Initialize services
	emailService := services.NewEmailService()

	// Initialize handlers
	emailHandler := rest.NewEmailHandler(emailService)

	// Setup routes
	app.Post("/send-email", emailHandler.SendEmail)

	// Initialize and start NATS handler
	natsHandler := messaging.NewNATSHandler(nc, emailService)
	natsHandler.Subscribe()

	// Start the application
	log.Fatal(app.Listen(":3002"))
}
