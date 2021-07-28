package main

import (
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/Toolnado/authorization-module/api"
	"github.com/Toolnado/authorization-module/internal/database"
	"github.com/Toolnado/authorization-module/internal/repository"
	"github.com/Toolnado/authorization-module/internal/rpc"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	var once sync.Once
	once.Do(func() {
		logrus.SetFormatter(new(logrus.JSONFormatter))

		if err := initConfig(); err != nil {
			logrus.Printf("error initializing config %s", err.Error())
		}

		if err := godotenv.Load(); err != nil {
			logrus.Printf("error loading env variables: %s", err.Error())
		}

	})

	start()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func start() {
	store := database.NewStore()
	repo := repository.NewRepository(store)
	grpcServer := grpc.NewServer()
	svr := rpc.NewGrpcServer(repo)
	api.RegisterAuthorizationServer(grpcServer, svr)
	listener, err := net.Listen("tcp", ":"+viper.GetString("port"))

	if err != nil {
		logrus.Printf("Error listen: %s", err.Error())
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			logrus.Printf("Error serve: %s", err.Error())
		}
	}()
	logrus.Printf("the server is running...")
	<-stopChan
	grpcServer.GracefulStop()
	logrus.Printf("server stopped...")
}
