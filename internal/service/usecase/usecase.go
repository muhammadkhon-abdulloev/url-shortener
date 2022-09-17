package usecase

import (
	"github.com/muhammadkhon-abdulloev/url-shortener/config"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/models"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/service"
)

type serviceUC struct {
	cfg *config.Config
}

func NewServiceUC(cfg *config.Config) service.UseCase {
	return &serviceUC{
		cfg: cfg,
	}
}

func (s *serviceUC) GetURL(id string) (response models.GetURLResponse, err error) {
	return
}

func (s *serviceUC) NewURL(params models.NewURLParams) (response models.NewURLResponse, err error) {
	return
}
