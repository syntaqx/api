package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syntaqx/api/internal/config"
	"github.com/syntaqx/api/internal/model"
)

//go:generate go run github.com/matryer/moq -pkg mock -out ./mock/weather_service.go . WeatherService

type WeatherService interface {
	GetWeather(location string) (*model.Weather, error)
}

type weatherService struct {
	apiKey string
}

// Assert weatherService implements WeatherService interface at comiple time.
var _ WeatherService = (*weatherService)(nil)

// NewWeatherService creates a new weather service.
func NewWeatherService(cfg *config.Config) *weatherService {
	return &weatherService{
		apiKey: cfg.WeatherAPIKey,
	}
}

type WeatherAPIResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempF     float64 `json:"temp_f"`
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		}
	} `json:"current"`
}

func (s *weatherService) GetWeather(location string) (*model.Weather, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", s.apiKey, location)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response WeatherAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &model.Weather{
		Location:    response.Location.Name,
		Temperature: fmt.Sprintf("%.1f°F", response.Current.TempF),
		Description: response.Current.Condition.Text,
	}, nil
}
