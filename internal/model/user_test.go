package model

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/render"
	"github.com/gofrs/uuid/v5"
)

func TestUser_Bind(t *testing.T) {
	user := &User{}
	// Create a mock request with JSON payload
	payload := `{"login": "test", "email": "test@example.com", "name": "Test User"}`
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	request.Header.Set("Content-Type", "application/json")

	err := render.Bind(request, user)
	if err != nil {
		t.Errorf("Bind method returned an error: %v", err)
	}

	// Verify the user fields are correctly populated
	if user.Login != "test" {
		t.Errorf("Expected login to be 'test', got '%s'", user.Login)
	}
	if user.Email != "test@example.com" {
		t.Errorf("Expected email to be 'test@example.com', got '%s'", user.Email)
	}
	if user.Name != "Test User" {
		t.Errorf("Expected name to be 'Test User', got '%s'", user.Name)
	}
}

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

func TestNewUserListResponse(t *testing.T) {
	users := []*User{
		{
			ID:    uuid.Nil,
			Login: "test",
			Email: "test",
			Name:  "test",
		},
	}

	list := NewUserListResponse(users)
	if len(list) != 1 {
		t.Errorf("Expected list length of 1, got %d", len(list))
	}
}
