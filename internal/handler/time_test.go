package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestTimeHandler_CurrentTime(t *testing.T) {
	handler := NewTimeHandler()

	req, err := http.NewRequest("GET", "/time", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	handler.RegisterRoutes(router)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response TimeResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedTime := time.Now().UTC().Format(time.RFC3339)
	assert.Equal(t, expectedTime, response.CurrentTime)
}
