package main

import (
	"shortener"
	"shortener/configs"
	"shortener/internal/handler"
	"shortener/internal/repository"
	"shortener/internal/service"
	"shortener/pkg/event"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(config)
	if err != nil {
		logrus.Fatalf("failed to initialization db: %s", err.Error())
	}

	eventBus := event.NewEventBus()

	repository := repository.NewRepository(db.DB)
	service := service.NewService(repository, config, eventBus)
	handler := handler.NewHandler(service)

	srv := new(shortener.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRouts()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
	logrus.Println("Application Started...")

}
