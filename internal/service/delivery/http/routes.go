package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/service"
)

func MapServiceRoutes(mx *chi.Mux, h service.Handlers) {
	mx.Get("/{id}", h.GetURL)
	mx.Post("/short/{url}", h.NewURL)
}
