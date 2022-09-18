package service

import (
	"net/http"
)

type Handlers interface {
	GetURL(w http.ResponseWriter, r *http.Request)
	NewURL(w http.ResponseWriter, r *http.Request)
}
