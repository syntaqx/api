package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/syntaqx/api/internal/service"
)

type WeatherHandler struct {
	service *service.WeatherService
}

func NewWeatherHandler(service *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) RegisterRoutes(r chi.Router) {
	r.Get("/weather", h.GetWeather)
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		location = "84095"
	}

	weather, err := h.service.GetWeather(location)
	if err != nil {
		http.Error(w, "failed to get weather", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, weather)
}
