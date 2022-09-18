package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/muhammadkhon-abdulloev/url-shortener/config"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/server"
	"github.com/muhammadkhon-abdulloev/url-shortener/pkg/logger"
)

func main() {
	cfg, err := config.ParseConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error while parsing config: %s", err.Error())
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()

	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s.", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	mx := chi.NewMux()

	srv := server.NewServer(cfg, &http.Server{
		Handler:        mx,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * cfg.Server.WriteTimeout,
	}, appLogger, mx)

	if err := srv.Run(); err != nil {
		appLogger.Fatalf("Error occured: %s", err.Error())
	}
}
