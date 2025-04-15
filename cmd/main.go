package main

import (
	"shortener"
	"shortener/internal/handler"

	"github.com/sirupsen/logrus"
)

func main() {
	handler := handler.NewHandler()
	srv := new(shortener.Server)
	if err := srv.Run("8888", handler.InitRouts()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
	logrus.Println("Application Started...")

}
