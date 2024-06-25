package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	id := uuid.Must(uuid.NewV4())
	// Create our base user
	createRequest := CreateUserRequest{
		Login:    "testuser",
		Email:    "testuser@example.com",
		Password: "testpassword",
		Name:     "Test User",
	}

	mockUserService := &mock.UserServiceMock{
		GetUserByIDFunc: func(id uuid.UUID) (*model.User, error) {
			return &model.User{
				ID:    id,
				Login: createRequest.Login,
				Email: createRequest.Email,
				Name:  createRequest.Name,
			}, nil
		},
		ListUsersFunc: func() ([]*model.User, error) {
			return []*model.User{
				{
					ID:    id,
					Login: createRequest.Login,
					Email: createRequest.Email,
					Name:  createRequest.Name,
				},
			}, nil
		},
		CreateUserFunc: func(user *model.User) error {
			user.ID = uuid.Must(uuid.NewV4())
			return nil
		},
		UpdateUserFunc: func(user *model.User) error {
			return nil
		},
		DeleteUserFunc: func(id uuid.UUID) error {
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

	// ===========================================================================
	// GET /users/{id}
	// ===========================================================================

	// Create a mock response recorder
	rr = httptest.NewRecorder()

	// Create a mock request
	req, err = http.NewRequest(http.MethodGet, UserURLPrefix+"/"+user.ID.String(), nil)
	assert.NoError(t, err)

	// Serve the request
	r.ServeHTTP(rr, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body
	err = json.NewDecoder(rr.Body).Decode(&user)
	assert.NoError(t, err)

	assert.NotEmpty(t, user.ID)
	assert.Equal(t, createRequest.Login, user.Login)
	assert.Equal(t, createRequest.Email, user.Email)
	assert.Equal(t, createRequest.Name, user.Name)

	// ===========================================================================
	// GET /users
	// ===========================================================================

	// Create a mock response recorder
	rr = httptest.NewRecorder()

	// Create a mock request
	req, err = http.NewRequest(http.MethodGet, UserURLPrefix, nil)
	assert.NoError(t, err)

	// Serve the request
	r.ServeHTTP(rr, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body
	var users []model.User
	err = json.NewDecoder(rr.Body).Decode(&users)
	assert.NoError(t, err)

	assert.NotEmpty(t, users)
	assert.Len(t, users, 1)

	// ===========================================================================
	// PUT /users/{id}
	// ===========================================================================

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

	// ===========================================================================
	// DELETE /users/{id}
	// ===========================================================================

	// Create a mock response recorder
	rr = httptest.NewRecorder()

	// Create a mock request
	req, err = http.NewRequest(http.MethodDelete, UserURLPrefix+"/"+user.ID.String(), nil)
	assert.NoError(t, err)

	// Serve the request
	r.ServeHTTP(rr, req)

	fmt.Printf("Status: %+v\n", rr.Code)
	fmt.Printf("Body: %+v\n", rr.Body.String())

	// Assert the response status code
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestUserHandler_ServiceErrors(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name           string
		endpoint       string
		method         string
		requestBody    interface{}
		expectedStatus int
	}{
		{
			name:           "CreateUser",
			endpoint:       UserURLPrefix,
			method:         http.MethodPost,
			requestBody:    CreateUserRequest{},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "GetUser",
			endpoint:       UserURLPrefix + "/00000000-0000-0000-0000-000000000000",
			method:         http.MethodGet,
			requestBody:    nil,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "ListUsers",
			endpoint:       UserURLPrefix,
			method:         http.MethodGet,
			requestBody:    nil,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "UpdateUser",
			endpoint:       UserURLPrefix + "/00000000-0000-0000-0000-000000000000",
			method:         http.MethodPut,
			requestBody:    model.User{},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "DeleteUser",
			endpoint:       UserURLPrefix + "/00000000-0000-0000-0000-000000000000",
			method:         http.MethodDelete,
			requestBody:    nil,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	// Create a mock service that returns errors
	mockErrorService := &mock.UserServiceMock{
		GetUserByIDFunc: func(id uuid.UUID) (*model.User, error) {
			return nil, errors.New("mock error")
		},
		ListUsersFunc: func() ([]*model.User, error) {
			return nil, errors.New("mock error")
		},
		CreateUserFunc: func(user *model.User) error {
			return errors.New("mock error")
		},
		UpdateUserFunc: func(user *model.User) error {
			return errors.New("mock error")
		},
		DeleteUserFunc: func(id uuid.UUID) error {
			return errors.New("mock error")
		},
	}

	// Create a user handler with the mock service
	h := NewUserHandler(mockErrorService)

	// Create a mock router
	r := chi.NewRouter()
	h.RegisterRoutes(r)

	// Iterate over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock response recorder
			rr := httptest.NewRecorder()

			// Create a mock request body
			requestBody, _ := json.Marshal(tc.requestBody)

			// Create a mock request
			req, err := http.NewRequest(tc.method, tc.endpoint, bytes.NewBuffer(requestBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			// Serve the request
			r.ServeHTTP(rr, req)

			// Assert the response status code
			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}
func TestUserHandler_BindErrors(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name           string
		endpoint       string
		method         string
		requestBody    interface{}
		expectedStatus int
	}{
		{
			name:           "CreateUser_BindError",
			endpoint:       UserURLPrefix,
			method:         http.MethodPost,
			requestBody:    "invalid_request_body",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "UpdateUser_BindError",
			endpoint:       UserURLPrefix + "/00000000-0000-0000-0000-000000000000",
			method:         http.MethodPut,
			requestBody:    "invalid_request_body",
			expectedStatus: http.StatusBadRequest,
		},
	}

	// Create a mock service
	mockUserService := &mock.UserServiceMock{}

	// Create a user handler with the mock service
	h := NewUserHandler(mockUserService)

	// Create a mock router
	r := chi.NewRouter()
	h.RegisterRoutes(r)

	// Iterate over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock response recorder
			rr := httptest.NewRecorder()

			// Create a mock request body
			requestBody, _ := json.Marshal(tc.requestBody)

			// Create a mock request
			req, err := http.NewRequest(tc.method, tc.endpoint, bytes.NewBuffer(requestBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			// Serve the request
			r.ServeHTTP(rr, req)

			// Assert the response status code
			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}
