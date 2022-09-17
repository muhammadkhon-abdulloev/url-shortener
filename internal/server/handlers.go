package server

import (
	"time"

	"github.com/go-chi/chi/v5/middleware"
	handlers "github.com/muhammadkhon-abdulloev/url-shortener/internal/service/delivery/http"
	usecase "github.com/muhammadkhon-abdulloev/url-shortener/internal/service/usecase"
)

func (s *Server) MapHandlers() {
	serviceUC := usecase.NewServiceUC(s.cfg)
	serviceHandler := handlers.NewServiceHandlers(s.cfg, serviceUC)

	s.mx.Use(middleware.Timeout(time.Second * 30))
	handlers.MapServiceRoutes(s.mx, serviceHandler)
}
