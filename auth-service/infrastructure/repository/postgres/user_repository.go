package infrastructure

import (
	"authService/domain"
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.db.NamedExecContext(ctx, `INSERT INTO users (username, password) VALUES (:username, :password)`, user)
	return err
}
