package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/muhammadkhon-abdulloev/url-shortener/config"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/models"
	"github.com/muhammadkhon-abdulloev/url-shortener/internal/service"
	"github.com/muhammadkhon-abdulloev/url-shortener/pkg/logger"
	"github.com/muhammadkhon-abdulloev/url-shortener/pkg/types"
)

type serviceHandles struct {
	cfg       *config.Config
	logger    logger.Logger
	serviceUC service.UseCase
}

func NewServiceHandlers(
	cfg *config.Config,
	logger logger.Logger,
	serviceUC service.UseCase,
) service.Handlers {
	return &serviceHandles{
		cfg:       cfg,
		logger:    logger,
		serviceUC: serviceUC,
	}
}

func (h *serviceHandles) GetLongURL(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	response, err := h.serviceUC.GetLongURL(id)
	if err != nil {
		h.writeError(w, err, http.StatusInternalServerError)
	}

	if response.Redirect {
		http.Redirect(w, r, response.RedirectURI, http.StatusSeeOther)
	}

	respBody, err := json.Marshal(response)
	if err != nil {
		h.writeError(w, err, http.StatusInternalServerError)
	}

	h.writeResponse(w, respBody)
}

func (h *serviceHandles) ShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.writeError(w, err, http.StatusBadRequest)
	}

	var params models.NewURLParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		h.writeError(w, err, http.StatusBadRequest)
	}
	response, err := h.serviceUC.ShortURL(params)
	if err != nil {
		h.writeError(w, err, http.StatusInternalServerError)
	}

	respBody, err := json.Marshal(response)
	if err != nil {
		h.writeError(w, err, http.StatusInternalServerError)
	}

	h.writeResponse(w, respBody)
}

func (h *serviceHandles) writeError(w http.ResponseWriter, err error, statusCode int) {
	h.logger.Error(err)
	http.Error(w, err.Error(), statusCode)
}

func (h *serviceHandles) writeResponse(w http.ResponseWriter, body []byte) {
	w.Header().Set(types.HeaderContentType, types.ApplicationJSON)
	w.Write(body)
}
