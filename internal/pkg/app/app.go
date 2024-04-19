package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/XRS0/Sigma-Network/configs"
	"github.com/XRS0/Sigma-Network/internal/app/handler"
	"github.com/XRS0/Sigma-Network/internal/app/repository"
	"github.com/XRS0/Sigma-Network/internal/app/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type App struct {
	db         *sqlx.DB
	handler    *handler.Handler
	service    *service.Service
	repository *repository.Repository
	server     *http.Server
}

func New() *App {
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

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	App := &App{
		db:         db,
		handler:    handler,
		service:    service,
		repository: repository,
	}

	return App
}

func (a *App) Run(port string) error {
	a.server = &http.Server{
		Addr:           ":" + port,
		Handler:        a.handler.InitRoutes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	if err := a.db.Close(); err != nil {
		log.Printf("[ERROR] failed to close db connection: %s", err.Error())
	}
	return a.server.Shutdown(ctx)
}
