package server

import (
	"fmt"
	"net/http"

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

func (s *Server) Run() (err error) {
	s.MapHandlers()

	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port), s.mx)
}
