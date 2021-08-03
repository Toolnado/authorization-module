package rpc

import (
	"context"
	"log"
	"time"

	"github.com/Toolnado/authorization-module/api"
	"github.com/Toolnado/authorization-module/internal/model"
	"github.com/Toolnado/authorization-module/internal/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "wefaef#$%1q2431543643rhGGESC24234235kqklw"
	tockenTTl  = 12 * time.Hour
)

type GrpcServer struct {
	api.UnimplementedAuthorizationServer
	Repository *repository.Repository
}

type tokenKlaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewGrpcServer(repo *repository.Repository) *GrpcServer {
	return &GrpcServer{
		Repository: repo,
	}
}

func (r *GrpcServer) SignUp(ctx context.Context, User *api.User) (*api.UserId, error) {
	user := conversion(User)
	userId, err := r.Repository.Authorization.CreateUser(user)

	if err != nil {
		log.Printf("Error create user: %s", err.Error())
		return nil, err
	}

	id := &api.UserId{
		Id: uint32(userId),
	}

	return id, nil
}

func (r *GrpcServer) SignIn(ctx context.Context, User *api.User) (*api.Token, error) {
	token := &api.Token{}
	user := conversion(User)
	targetUser, err := r.Repository.Authorization.GetUser(user.Username, User.Password)
	if err != nil {
		return token, err
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenKlaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tockenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		targetUser.Id,
	})

	token.Token, err = newToken.SignedString([]byte(signingKey))
	if err != nil {
		return token, err
	}

	return token, nil

}

func conversion(user *api.User) *model.User {
	newUser := &model.User{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	}

	return newUser
}
