package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Save the current value of the PORT environment variable
	oldPort := os.Getenv("PORT")
	// Set a temporary value for the PORT environment variable
	os.Setenv("PORT", "8080")
	// Restore the original value of the PORT environment variable when the test is done
	defer os.Setenv("PORT", oldPort)

	// Create a new request to the root URL
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Call the main function, which starts the HTTP server
	go main()

	// Wait for the server to start
	time.Sleep(time.Second)

	// Send the request to the server
	http.DefaultServeMux.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

}
