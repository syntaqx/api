package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TimeHandler struct {
}

func NewTimeHandler() *TimeHandler {
	return &TimeHandler{}
}

func (h *TimeHandler) RegisterRoutes(r chi.Router) {
	r.Get("/time", h.CurrentTime)
}

type TimeResponse struct {
	CurrentTime string `json:"currentTime"`
}

func (resp *TimeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *TimeHandler) CurrentTime(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Render(w, r, &TimeResponse{CurrentTime: time.Now().UTC().Format(time.RFC3339)})
}
