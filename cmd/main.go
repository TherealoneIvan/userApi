package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	userapi "refactoring"
	"refactoring/api"
	"refactoring/pkg/handler"
	"refactoring/pkg/repository"
)

const store = `users.json`

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error occured running the configs %s", err.Error())
	}
	repos := repository.NewRepository(store)
	userApi := api.NewUserApiService(repos)
	handlers := handler.NewHandler(userApi)
	server := new(userapi.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured running the server %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
