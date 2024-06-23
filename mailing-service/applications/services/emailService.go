package services

import (
	"fmt"
	"log"
	"mailingService/domain"
	"net/smtp"
	"os"
)

type EmailService struct {
	SMTPHost string
	SMTPPort string
}

func NewEmailService() *EmailService {
	return &EmailService{
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: os.Getenv("SMTP_PORT"),
	}
}

func (s *EmailService) SendWelcomeEmail(emailReq *domain.EmailRequest) error {
	log.Printf("Sending welcome email to %s at %s hosted on %s and serving on port %s", emailReq.Username, emailReq.Email, s.SMTPHost, s.SMTPPort)

	from := "no-reply@ecommerce.com"
	to := emailReq.Email
	subject := "Welcome to E-commerce!"
	body := fmt.Sprintf("Hello %s,\n\nThank you for signing up. We're excited to have you with us. \n\nCheers from all of us at BetMate", emailReq.Username)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	addr := fmt.Sprintf("%s:%s", s.SMTPHost, s.SMTPPort)

	err := smtp.SendMail(addr, nil, from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Printf("Welcome email sent to %s", to)
	return nil
}
