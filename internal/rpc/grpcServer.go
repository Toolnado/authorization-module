package rpc

import (
	"context"

	"github.com/Toolnado/authorization-module/api"
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
	return nil, nil
}
