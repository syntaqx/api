package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestRootHandler_Root(t *testing.T) {
	handler := NewRootHandler()
	router := chi.NewRouter()
	handler.RegisterRoutes(router)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
