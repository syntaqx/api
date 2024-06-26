package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestBookmarksCRUD(t *testing.T) {
	// Create a new router
	r := chi.NewRouter()

	// Create a new GamesHandler
	handler := NewBookmarksHandler()

	// Register routes
	handler.RegisterRoutes(r)

	// Create a new request
	req, err := http.NewRequest(http.MethodGet, BookmarksURLPrefix, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// TODO: Add assertions for the response body and any other CRUD operations

	// Example assertion for the response body
	// assert.Equal(t, expectedResponseBody, rr.Body.String())
}
