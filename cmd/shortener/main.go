package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/muhammadkhon-abdulloev/url-shortener/config"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/server"
)

func main() {
	cfg, err := config.ParseConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error while parsing config: %s", err.Error())
	}

	mx := chi.NewMux()

	srv := server.NewServer(cfg, mx)

	if err := srv.Run(); err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}
