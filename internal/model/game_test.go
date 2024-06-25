package model

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/render"
	"github.com/gofrs/uuid/v5"
)

func TestGame_Bind(t *testing.T) {
	game := &Game{}
	// Create a mock request with JSON payload
	payload := `{"name": "Test Game"}`
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	request.Header.Set("Content-Type", "application/json")

	err := render.Bind(request, game)
	if err != nil {
		t.Errorf("Bind method returned an error: %v", err)
	}

	// Verify the game fields are correctly populated
	if game.Name != "Test Game" {
		t.Errorf("Expected name to be 'Test Game', got '%s'", game.Name)
	}
}

func TestGame_Render(t *testing.T) {
	game := &Game{}

	// Create a mock response writer and request
	responseWriter := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	err := game.Render(responseWriter, request)
	if err != nil {
		t.Errorf("Render method returned an error: %v", err)
	}
}

func TestewGameListResponse(t *testing.T) {
	games := []*Game{
		{
			ID:   uuid.Nil,
			Name: "test",
		},
	}

	list := NewGameListResponse(games)
	if len(list) != 1 {
		t.Errorf("Expected list to have 1 element, got %d", len(list))
	}
}
