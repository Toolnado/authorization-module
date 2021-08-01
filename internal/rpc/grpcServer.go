package rpc

import (
	"context"
	"log"

	"github.com/Toolnado/authorization-module/api"
	"github.com/Toolnado/authorization-module/internal/model"
	"github.com/Toolnado/authorization-module/internal/repository"
)

type GrpcServer struct {
	api.UnimplementedAuthorizationServer
	Repository *repository.Repository
}

func NewGrpcServer(repo *repository.Repository) *GrpcServer {
	return &GrpcServer{
		Repository: repo,
	}
}

func (r *GrpcServer) CreateUser(ctx context.Context, User *api.User) (*api.UserId, error) {
	user := conversion(User)
	userId, err := r.Repository.Authorization.CreateUser(ctx, user)

	if err != nil {
		log.Printf("Error create user: %s", err.Error())
		return nil, err
	}

	id := &api.UserId{
		Id: uint32(userId),
	}

	return id, nil
}

func conversion(user *api.User) *model.User {
	newUser := &model.User{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	}

	return newUser
}
