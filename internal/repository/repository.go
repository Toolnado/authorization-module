package repository

import (
	"context"

	"github.com/Toolnado/authorization-module/internal/model"
)

type Authorization interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
}

type Repository struct {
	Authorization Authorization
}

func NewRepository(auth Authorization) *Repository {
	return &Repository{
		Authorization: auth,
	}
}
