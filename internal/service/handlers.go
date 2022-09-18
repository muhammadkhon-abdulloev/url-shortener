package service

import (
	"net/http"
)

type Handlers interface {
	GetLongURL(w http.ResponseWriter, r *http.Request)
	ShortURL(w http.ResponseWriter, r *http.Request)
}
