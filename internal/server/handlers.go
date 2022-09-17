package server

import (
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) MapHandlers() {
	s.mx.Use(middleware.Timeout(time.Second * 30))

	
}