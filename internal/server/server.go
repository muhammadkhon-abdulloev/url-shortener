package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/muhammadkhon-abdulloev/url-shortener/config"
)

type Server struct {
	cfg *config.Config
	mx  *chi.Mux
}

// NewServer - constructor function.
func NewServer(cfg *config.Config, mx *chi.Mux) *Server {
	return &Server{
		cfg: cfg,
		mx:  mx,
	}
}

func (s *Server) Server() {
	
}
