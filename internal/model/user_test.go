package model

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUser_Render(t *testing.T) {
	user := &User{}

	// Create a mock response writer and request
	responseWriter := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	err := user.Render(responseWriter, request)
	if err != nil {
		t.Errorf("Render method returned an error: %v", err)
	}
}
