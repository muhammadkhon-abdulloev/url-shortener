package service

import (
	"net/http"
)

type Handlers interface {
	GetURL(w http.ResponseWriter, r *http.Request)
	ShortURL(w http.ResponseWriter, r *http.Request)
}
