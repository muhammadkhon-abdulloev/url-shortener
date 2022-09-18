package service

import "github.com/muhammadkhon-abdulloev/url-shortener/internal/models"

type Repository interface {
	GetLongURL(id string) (response *models.GetURLResponse, err error)
	ShortURL(params models.NewURLParams) (response *models.NewURLResponse, err error)
}
