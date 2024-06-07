package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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

func (h *TimeHandler) CurrentTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	response := TimeResponse{CurrentTime: currentTime}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
