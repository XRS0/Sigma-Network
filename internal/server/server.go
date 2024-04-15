package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context, db *sqlx.DB) error {
	if err := db.Close(); err != nil {
		log.Printf("[ERROR] failed to close db connection: %s", err.Error())
	}
	return s.httpServer.Shutdown(ctx)
}