package main

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	var once sync.Once
	once.Do(func() {
		logrus.SetFormatter(new(logrus.JSONFormatter))

		if err := initConfig(); err != nil {
			logrus.Fatalf("error initializing config %s", err.Error())
		}

		if err := godotenv.Load(); err != nil {
			logrus.Fatalf("error loading env variables: %s", err.Error())
		}

	})

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
