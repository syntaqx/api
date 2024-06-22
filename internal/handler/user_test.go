package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/assert"

	"github.com/syntaqx/api/internal/model"
	"github.com/syntaqx/api/internal/service/mock"
)

func TestUserHandler_CRUD(t *testing.T) {
	// Create our base user
	createRequest := CreateUserRequest{
		Login:    "testuser",
		Email:    "testuser@example.com",
		Password: "testpassword",
		Name:     "Test User",
	}

	mockUserService := &mock.UserServiceMock{
		CreateUserFunc: func(user *model.User) error {
			user.ID = uuid.Must(uuid.NewV4())
			return nil
		},
		UpdateUserFunc: func(user *model.User) error {
			return nil
		},
	}

	h := NewUserHandler(mockUserService)

	// Create a mock response recorder
	rr := httptest.NewRecorder()

	// Create a mock router
	r := chi.NewRouter()
	h.RegisterRoutes(r)

	// Create a mock request body
	requestBody, _ := json.Marshal(createRequest)

	// Create a mock request
	req, err := http.NewRequest(http.MethodPost, UserURLPrefix, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	// Serve the request
	r.ServeHTTP(rr, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body
	var user model.User
	err = json.NewDecoder(rr.Body).Decode(&user)
	assert.NoError(t, err)

	assert.NotEmpty(t, user.ID)
	assert.Equal(t, createRequest.Login, user.Login)
	assert.Equal(t, createRequest.Email, user.Email)
	assert.Equal(t, createRequest.Name, user.Name)

	// Update the user
	updateRequest := model.User{
		Login: user.Login,
		Email: user.Email,
		Name:  "Updated Name",
	}

	// Create a mock response recorder
	rr = httptest.NewRecorder()

	// Create a mock request body
	requestBody, _ = json.Marshal(updateRequest)

	// Create a mock request
	req, err = http.NewRequest(http.MethodPut, UserURLPrefix+"/"+user.ID.String(), bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	// Serve the request
	r.ServeHTTP(rr, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body
	err = json.NewDecoder(rr.Body).Decode(&user)
	assert.NoError(t, err)

	assert.NotEmpty(t, user.ID)
	assert.Equal(t, updateRequest.Login, user.Login)
	assert.Equal(t, updateRequest.Email, user.Email)
	assert.Equal(t, updateRequest.Name, user.Name)
}
