package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/XRS0/Sigma-Network/configs"
	"github.com/XRS0/Sigma-Network/internal/handler"
	"github.com/XRS0/Sigma-Network/internal/repository"
	"github.com/XRS0/Sigma-Network/internal/server"
	"github.com/XRS0/Sigma-Network/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("[ERROR] failed to initialize configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("[ERROR] failed to load env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDb()
	if err != nil {
		log.Fatalf("[ERROR] failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)
	server := new(server.Server)
	
	go func() {
		server.Run(viper.GetString("port"), handler.InitRoutes())
	}()

	log.Print("Sigma-Network started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	log.Print("Sigma-Network Shutting Down")

	if err := server.Shutdown(context.Background(), db); err != nil {
		log.Printf("[ERROR] failed to shut down server: %s", err.Error())
	}
}