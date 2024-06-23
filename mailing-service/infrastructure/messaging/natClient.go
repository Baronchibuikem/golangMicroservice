package messaging

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func NewNATSClient() *nats.Conn {
	natsUrl := os.Getenv("NATS_URL")
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Printf("Cannot connect to NATS with URL %s on mailingService because of %s", natsUrl, err)
	}
	log.Printf("Successfully connected to NATS with URL %s on mailingService", natsUrl)
	return nc
}
