package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/syntaqx/api/internal/model"
	"github.com/syntaqx/api/internal/service/mock"
)

func TestWeatherHandler_GetWeather(t *testing.T) {
	mockResponse := &model.Weather{
		Location:    "test",
		Temperature: "test",
		Description: "test",
	}

	mockedWeatherService := &mock.WeatherServiceMock{
		GetWeatherFunc: func(location string) (*model.Weather, error) {
			return mockResponse, nil
		},
	}

	// Create a new weather handler with the mocked service
	weatherHandler := NewWeatherHandler(mockedWeatherService)

	// Create a new router and register the weather handler
	router := chi.NewRouter()
	weatherHandler.RegisterRoutes(router)

	// Create a new request with a location query parameter
	req, err := http.NewRequest("GET", "/weather?location=test", nil)
	assert.NoError(t, err)

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Assert the response status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestWeatherHandler_GetWeatherDefaultLocation(t *testing.T) {
	mockResponse := &model.Weather{
		Location:    "test",
		Temperature: "test",
		Description: "test",
	}

	mockedWeatherService := &mock.WeatherServiceMock{
		GetWeatherFunc: func(location string) (*model.Weather, error) {
			return mockResponse, nil
		},
	}

	// Create a new weather handler with the mocked service
	weatherHandler := NewWeatherHandler(mockedWeatherService)

	// Create a new router and register the weather handler
	router := chi.NewRouter()
	weatherHandler.RegisterRoutes(router)

	// Create a new request with a location query parameter
	req, err := http.NewRequest("GET", "/weather", nil)
	assert.NoError(t, err)

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Assert the response status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestWeatherHandler_GetWeatherError(t *testing.T) {
	mockedWeatherService := &mock.WeatherServiceMock{
		GetWeatherFunc: func(location string) (*model.Weather, error) {
			return nil, errors.New("example error")
		},
	}

	// Create a new weather handler with the mocked service
	weatherHandler := NewWeatherHandler(mockedWeatherService)

	// Create a new router and register the weather handler
	router := chi.NewRouter()
	weatherHandler.RegisterRoutes(router)

	// Create a new request with a location query parameter
	req, err := http.NewRequest("GET", "/weather?location=test", nil)
	assert.NoError(t, err)

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Assert the response status code is OK
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
