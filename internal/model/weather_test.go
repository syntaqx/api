package model

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWeather_Render(t *testing.T) {
	weather := &Weather{} // Create a new Weather instance

	// Create a mock response writer and request
	responseWriter := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	err := weather.Render(responseWriter, request)
	if err != nil {
		t.Errorf("Render method returned an error: %v", err)
	}
}
