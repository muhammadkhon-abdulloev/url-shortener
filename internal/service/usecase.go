package service

import "github.com/muhammadkhon-abdulloev/url-shortener/internal/models"

type UseCase interface {
	GetURL(id string) (response models.GetURLResponse, err error)
	NewURL(params models.NewURLParams) (response models.NewURLResponse, err error)
}
