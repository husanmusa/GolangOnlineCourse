package main

import (
	"auth/config"
	"auth/pkg/app"
	"auth/pkg/handler"
	"auth/pkg/pubsub"
	"auth/pkg/repository"
	"auth/pkg/repository/postgres"
	"auth/pkg/service"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(".env"); err != nil {
		slog.Error(fmt.Sprintf("could not read .env file %s", err.Error()))
	}

	if err := config.InitConfig(); err != nil {
		slog.Error(err.Error())
	}

	postgresDB, err := postgres.NewPostgres(postgres.PostgresConfig{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		User:     viper.GetString("postgres.user"),
		Password: "pass",
		Dbname:   viper.GetString("postgres.dbName"),
		SSLMode:  viper.GetString("postgres.sslMode"),
	})

	if err != nil {
		slog.Error(fmt.Sprintf("postgres: %s", err.Error()))
	}
	repositories := repository.NewRepository(postgresDB)
	ps := pubsub.NewPubSub()
	services := service.NewService(repositories, ps)
	h := handler.NewHandler(services, ps)

	server := app.Server{RPC: h}

	err = server.Run()
	if err != nil {
		slog.Error(fmt.Sprintf("error when initialize server: %s", err.Error()))
	}
}
