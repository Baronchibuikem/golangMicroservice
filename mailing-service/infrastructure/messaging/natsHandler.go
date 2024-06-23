package messaging

import (
	"encoding/json"
	"log"
	"mailingService/applications/services"
	"mailingService/domain"

	"github.com/nats-io/nats.go"
)

type NATSHandler struct {
	nc           *nats.Conn
	emailService *services.EmailService
}

func NewNATSHandler(nc *nats.Conn, emailService *services.EmailService) *NATSHandler {
	return &NATSHandler{nc: nc, emailService: emailService}
}

func (h *NATSHandler) Subscribe() {
	h.nc.Subscribe("user.created", h.handleUserCreatedEvent)
}

func (h *NATSHandler) handleUserCreatedEvent(msg *nats.Msg) {
	var user domain.User
	if err := json.Unmarshal(msg.Data, &user); err != nil {
		log.Printf("Error unmarshaling user.created event: %v", err)
		return
	}

	log.Printf("Received user.created event for user: %s", user.Username)

	emailReq := &domain.EmailRequest{
		Username: user.Username,
		Email:    user.Username + "@example.com",
	}

	// Send email directly
	err := h.emailService.SendWelcomeEmail(emailReq)
	if err != nil {
		log.Printf("Failed to send welcome email: %v", err)
	}
}
