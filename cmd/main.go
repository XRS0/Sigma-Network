package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/XRS0/Sigma-Network/internal/pkg/app"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	app := app.New()

	go func() {
		app.Run(viper.GetString("port"))
	}()

	log.Print("[INFO] Sigma-Network started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := app.Shutdown(context.Background()); err != nil {
		log.Printf("[ERROR] failed to shut down server: %s", err.Error())
	}

	log.Print("[INFO] Sigma-Network exited")
}