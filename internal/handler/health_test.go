package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler_GetHealth(t *testing.T) {
	handler := NewHealthHandler()

	router := chi.NewRouter()
	handler.RegisterRoutes(router)

	req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetHealth(rr, req)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedResponse := &HealthResponse{Status: http.StatusText(http.StatusOK)}
	actualResponse := &HealthResponse{}
	err = json.NewDecoder(rr.Body).Decode(actualResponse)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestHealthHandler_GetReady(t *testing.T) {
	handler := NewHealthHandler()

	router := chi.NewRouter()
	handler.RegisterRoutes(router)

	req, err := http.NewRequest(http.MethodGet, "/readiness", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetReady(rr, req)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedResponse := &ReadyResponse{Status: http.StatusText(http.StatusOK)}
	actualResponse := &ReadyResponse{}
	err = json.NewDecoder(rr.Body).Decode(actualResponse)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
