package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syntaqx/api/internal/config"
)

func TestGetWeather(t *testing.T) {
	weatherResponse := WeatherAPIResponse{
		Location: WeatherAPIResponseLocation{
			Name: "Test Location",
		},
		Current: WeatherAPIResponseCurrent{
			TempF: 75.0,
			TempC: 23.9,
			Condition: WeatherAPIResponseCurrentCondition{
				Text: "Sunny",
			},
		},
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(weatherResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))
	defer svr.Close()

	weatherService := NewWeatherService(&config.Config{
		WeatherAPIHost: svr.URL,
		WeatherAPIKey:  "test",
	})

	weather, err := weatherService.GetWeather(weatherResponse.Location.Name)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, weatherResponse.Location.Name, weather.Location)
	assert.Equal(t, fmt.Sprintf("%.1f°F", weatherResponse.Current.TempF), weather.Temperature)
	assert.Equal(t, weatherResponse.Current.Condition.Text, weather.Description)
}

func TestGetWeather_Fail(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "error", http.StatusInternalServerError)
	}))
	defer srv.Close()

	weatherService := NewWeatherService(&config.Config{
		WeatherAPIHost: srv.URL,
		WeatherAPIKey:  "test",
	})

	_, err := weatherService.GetWeather("Test Location")
	assert.Error(t, err)
}

func TestGetWeather_InvalidResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("invalid json"))
	}))
	defer srv.Close()

	weatherService := NewWeatherService(&config.Config{
		WeatherAPIHost: srv.URL,
		WeatherAPIKey:  "test",
	})

	_, err := weatherService.GetWeather("Test Location")
	assert.Error(t, err)
}
