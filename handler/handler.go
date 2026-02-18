package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"url-shortener/service"
)

type Handler struct {
	service *service.URLService
}

func NewHandler(service *service.URLService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL string `json:"url"`
	}

	_ = json.NewDecoder(r.Body).Decode(&request)

	code := h.service.Shorten(request.URL)

	response := map[string]string{
		"short_url": "http://localhost:8080/" + code,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")

	url, err := h.service.GetOriginal(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
