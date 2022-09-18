package usecase

import (
	"database/sql"
	"errors"
	"math/rand"
	"time"

	"github.com/muhammadkhon-abdulloev/url-shortener/config"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/models"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/service"
	"github.com/muhammadkhon-abdulloev/url-shortener/pkg/generator"
	"github.com/muhammadkhon-abdulloev/url-shortener/pkg/logger"
	"github.com/muhammadkhon-abdulloev/url-shortener/pkg/types"
)

type serviceUC struct {
	cfg    *config.Config
	logger logger.Logger
	repo   service.Repository
}

func NewServiceUC(
	cfg *config.Config,
	logger logger.Logger,
	repo service.Repository,
) service.UseCase {
	return &serviceUC{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
	}
}

func (s *serviceUC) GetLongURL(id string) (response *models.GetURLResponse, err error) {
	response, err = s.repo.GetLongURL(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("url not found")
		}
		s.logger.Errorf("Error occured while getting url: %s", err.Error())
		return nil, types.ErrOccured
	}
	response.Comment = "OK"
	response.Redirect = true
	response.ShortURI = s.cfg.Server.BaseURL + id

	return
}

func (s *serviceUC) ShortURL(params models.NewURLParams) (*models.NewURLResponse, error) {
	response := &models.NewURLResponse{}
	rand.Seed(time.Now().UnixMilli())
	id := generator.NewID(rand.Intn(s.cfg.Server.MaxVal))
	params.Shortened = id
	resp, err := s.repo.GetLongURL(id)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, types.ErrOccured
		}

		response, err = s.repo.ShortURL(params)
		if err != nil {
			return nil, types.ErrOccured
		}

		response.Message = "OK"
		response.Shortened = s.cfg.Server.BaseURL + response.Shortened

		return response, nil
	}

	if resp.RedirectURI == params.URL {
		response.Message = "OK"
		response.Shortened = s.cfg.Server.BaseURL + id

		return response, nil
	}

	return s.ShortURL(params)
}
