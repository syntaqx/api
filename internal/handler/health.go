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

// GetHealth godoc
// @Summary      Health check
// @Description  get the current service health
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  HealthResponse
// @Router       /healthz [get]
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

// GetHealth godoc
// @Summary      Readiness check
// @Description  get the current service readiness
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  ReadyResponse
// @Router       /readiness [get]
func (h *HealthHandler) GetReady(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Render(w, r, &ReadyResponse{Status: http.StatusText(http.StatusOK)})
}
