package model

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid/v5"
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
