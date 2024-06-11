package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) RegisterRoutes(r chi.Router) {
	r.Get("/healthz", h.GetHealth)
	r.Get("/readiness", h.GetReady)
}

type HealthResponse struct {
	Status string `json:"status"`
}

func (resp *HealthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Render(w, r, &HealthResponse{Status: http.StatusText(http.StatusOK)})
}

type ReadyResponse struct {
	Status string `json:"status"`
}

func (resp *ReadyResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *HealthHandler) GetReady(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Render(w, r, &ReadyResponse{Status: http.StatusText(http.StatusOK)})
}
