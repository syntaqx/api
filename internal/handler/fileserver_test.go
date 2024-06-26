package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestFileServer_FileNotFound(t *testing.T) {
	r := chi.NewRouter()
	root := http.Dir(".")

	// Call the FileServer function
	FileServer(r, "/static", root)

	// Create a new request to the static file URL
	req, err := http.NewRequest(http.MethodGet, "/static/file-not-found", nil)
	assert.NoError(t, err)

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestFileServer_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	r := chi.NewRouter()
	FileServer(r, "/*", http.Dir("."))
}
