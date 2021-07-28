package repository

import (
	"context"

	"github.com/Toolnado/authorization-module/internal/model"
)

type Authorization interface {
	CreateUser(ctx context.Context, user *model.User) (string, error)
}

type Repository struct {
	Authorization Authorization
}

func NewRepository(auth Authorization) *Repository {
	return &Repository{
		Authorization: auth,
	}
}
