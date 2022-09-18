package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/models"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/service"
)

type serviceRepo struct {
	db *sqlx.DB
}

func NewServiceRepo(db *sqlx.DB) service.Repository {
	return &serviceRepo{db: db}
}

func (r *serviceRepo) GetLongURL(id string) (*models.GetURLResponse, error) {
	var response models.GetURLResponse
	err := r.db.Get(&response.RedirectURI, queryGetLongURL, id)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *serviceRepo) ShortURL(params models.NewURLParams) (*models.NewURLResponse, error) {
	var response models.NewURLResponse
	err := r.db.Get(&response.Shortened, queryIsExist, params.URL)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		err = r.db.Get(&response.Shortened, queryNewShortened, params.Shortened, params.URL)
		if err != nil {
			return nil, err
		}
	}

	return &response, nil
}
