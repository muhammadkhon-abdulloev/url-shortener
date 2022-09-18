package usecase

import (
	"github.com/muhammadkhon-abdulloev/url-shortener/config"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/models"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/service"
	"github.com/muhammadkhon-abdulloev/url-shortener/pkg/logger"
)

type serviceUC struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewServiceUC(cfg *config.Config, logger logger.Logger) service.UseCase {
	return &serviceUC{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *serviceUC) GetURL(id string) (response models.GetURLResponse, err error) {
	return
}

func (s *serviceUC) NewURL(params models.NewURLParams) (response models.NewURLResponse, err error) {
	return
}
