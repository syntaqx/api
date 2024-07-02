package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"

	"github.com/syntaqx/api/internal/config"
	"github.com/syntaqx/api/internal/router"
)

func TestNewRouter(t *testing.T) {
	cfg := &config.Config{
		FQDN: "localhost",
		Port: "8080",
	}

	logger := zap.NewNop()

	r := router.NewRouter(cfg, logger)

	// Create a mock request
	request := httptest.NewRequest(http.MethodGet, "/swagger/doc.json", nil)

	// Create a mock response writer
	responseWriter := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(responseWriter, request)

	// Check the response status code
	if responseWriter.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseWriter.Code)
	}
}
