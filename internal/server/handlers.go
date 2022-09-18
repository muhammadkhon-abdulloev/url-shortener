package server

import (
	"time"

	"github.com/go-chi/chi/v5/middleware"
	handlers "github.com/muhammadkhon-abdulloev/url-shortener/internal/service/delivery/http"
	repository "github.com/muhammadkhon-abdulloev/url-shortener/internal/service/repository"
	usecase "github.com/muhammadkhon-abdulloev/url-shortener/internal/service/usecase"
)

func (s *Server) MapHandlers() {
	serviceRepo := repository.NewServiceRepo(s.db)
	serviceUC := usecase.NewServiceUC(s.cfg, s.logger, serviceRepo)
	serviceHandler := handlers.NewServiceHandlers(s.cfg, s.logger, serviceUC)

	s.mx.Use(middleware.Timeout(time.Second * 30))
	handlers.MapServiceRoutes(s.mx, serviceHandler)
}
