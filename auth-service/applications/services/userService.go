package applications

import (
	"authService/domain"
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
}

type UserService struct {
	repo UserRepository
	nc   *nats.Conn
}

func NewUserService(repo UserRepository, nc *nats.Conn) *UserService {
	return &UserService{repo: repo, nc: nc}
}

func (s *UserService) RegisterUser(ctx context.Context, user *domain.User) error {
	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = s.nc.Publish("user.created", userData)
	if err != nil {
		return err
	}
	log.Println("User has been created and published on nats server successfully.")
	return nil
}
