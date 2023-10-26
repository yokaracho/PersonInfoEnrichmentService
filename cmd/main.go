package main

import (
	peopleService "NameService"
	Nhandler "NameService/pkg/handler"
	logger "NameService/pkg/logger"
	postgres "NameService/pkg/repository/postgres"
	Nservice "NameService/pkg/service"
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	getLogger := logger.GetLogger()

	if err := godotenv.Load("/home/dev/NameService/.env"); err != nil {
		getLogger.Fatalf("error loading env variables: %s", err.Error())
	}
	getLogger.Infof("env variables successfully loaded")
	postgresCfg := postgres.Config{
		Username: os.Getenv("Username"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("POST_PORT"),
		DBName:   os.Getenv("DBName"),
	}

	postgresPool, err := postgres.NewConnectionPool(context.Background(), postgresCfg)
	if err != nil {
		getLogger.Fatalf("database connection error: %s", err.Error())
	}
	getLogger.Infof("База данных подключена успешно")
	getLogger.Debugf("База данных подключена успешно")

	repository := postgres.NewRepository(postgresPool)
	service := Nservice.NewService(repository)
	handler := Nhandler.NewHandler(service)

	server := new(peopleService.Server)

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			logrus.Infof("server shutdown error: %s", err.Error())
			logrus.Debugf("server shutdown error: %s", err.Error())
		}
	}()

	if err := server.Run(os.Getenv("PORT"), handler.GetRouter()); err != nil {
		log.Fatalf("server running error: %s", err.Error())
	} else {
		logrus.Infof("Server is running on port %s", os.Getenv("PORT"))
		logrus.Debugf("Server is running on port %s", os.Getenv("PORT"))
	}
}
