package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/muhammadkhon-abdulloev/url-shortener/config"
	"github.com/muhammadkhon-abdulloev/url-shortener/pkg/logger"
)

type Server struct {
	cfg        *config.Config
	logger     logger.Logger
	httpServer *http.Server
	mx         *chi.Mux
}

// NewServer - constructor function.
func NewServer(
	cfg *config.Config,
	httpServer *http.Server,
	logger logger.Logger,
	mx *chi.Mux,
) *Server {
	if cfg.Server.Mode == "Developement" {
		httpServer.Addr = cfg.Server.DevPort

	} else {
		httpServer.Addr = cfg.Server.Port
	}
	return &Server{
		cfg:        cfg,
		httpServer: httpServer,
		logger:     logger,
		mx:         mx,
	}
}

//  Run method which starts server and returns error if it occured.
func (s *Server) Run() (err error) {
	s.MapHandlers()
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			s.logger.Fatalf("Error occured while starting server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*60)
	defer shutdown()

	s.logger.Info("Server exited properly")
	return s.httpServer.Shutdown(ctx)
}
