package main

import (
	"farmish"
	"farmish/config"
	"farmish/constants"
	"farmish/pkg/handler"
	"farmish/pkg/pubsub"
	"farmish/pkg/repository"
	"farmish/pkg/repository/postgres"
	"farmish/pkg/service"
	pb "farmish/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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

	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthServiceClient(conn)

	h := handler.NewHandler(services, ps, c)

	routes := h.InitRoutes()

	server := new(farmish.Server)
	err = server.Run(viper.GetString("server.port"), routes)

	if err != nil {
		slog.Error(fmt.Sprintf("error when initialize server: %s", err.Error()))
	}

	go services.Animal.MonitorHealth(constants.HEALTH_MONITORING_TIMEOUT)
	go services.Animal.MonitorFeed(constants.FEED_MONITORING_TIMEOUT)
	h.ListenNotifications()
}
