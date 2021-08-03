package repository

import (
	"github.com/Toolnado/authorization-module/internal/model"
)

type Authorization interface {
	CreateUser(user *model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Repository struct {
	Authorization Authorization
}

func NewRepository(auth Authorization) *Repository {
	return &Repository{
		Authorization: auth,
	}
}
