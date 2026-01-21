package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nelsonmarro/vapulabs-web/config"
	"github.com/nelsonmarro/vapulabs-web/internal/web/handlers"
	"github.com/nelsonmarro/vapulabs-web/internal/web/services"
)

type Server struct {
	Config   *config.Config
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewServer(cfg *config.Config, infoLog, errorLog *log.Logger) *Server {
	return &Server{
		Config:   cfg,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}

func (s *Server) Serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.Config.Port),
		Handler:           s.getRoutes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	s.InfoLog.Printf("Starting server on port %d", s.Config.Port)

	return srv.ListenAndServe()
}

func (s *Server) createContactHandler() *handlers.ContactHandler {
	emailService := services.NewEmailService(s.Config)
	return handlers.NewContactHandler(emailService, s.ErrorLog)
}