package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/syntaqx/api/internal/service"
)

// DefaultLocation is the default location to use if none is provided.
const DefaultLocation = "84095"

const (
	WeatherURLPrefix = "/weather"
)

type WeatherHandler struct {
	service service.WeatherService
}

func NewWeatherHandler(service service.WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) RegisterRoutes(r chi.Router) {
	r.Get(WeatherURLPrefix, h.GetWeather)
}

// @Summary      Get the current weather
// @Description  Get the current weather
// @Router       /weater [get]
// @Tags         weather
// @Accept       json
// @Produce      json
// @Param        location   query  string  false  "Location Name"
// @Success      200  {object}  model.Weather
func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		location = DefaultLocation
	}

	weather, err := h.service.GetWeather(location)
	if err != nil {
		http.Error(w, "failed to get weather", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, weather)
}
