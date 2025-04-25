package main

import (
	"context"
	"os"
	"os/signal"
	"shortly"
	"shortly/configs"
	"shortly/internal/handler"
	"shortly/internal/repository"
	"shortly/internal/service"
	"shortly/pkg/event"
	"syscall"

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

	repository := repository.NewRepository(db)
	service := service.NewService(repository, config, eventBus)
	handler := handler.NewHandler(service)

	srv := new(shortly.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRouts()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()
	logrus.Println("Application Started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Application Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorf("error occurred on db connection close: %s", err.Error())
	}
	sqlDB.Close()
}
