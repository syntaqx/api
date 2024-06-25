package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid/v5"

	"github.com/syntaqx/api/internal/model"
	"github.com/syntaqx/api/internal/service"
)

const (
	UserURLPrefix = "/users"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(r chi.Router) {
	r.Route(UserURLPrefix, func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Get("/{id}", h.GetUser)
		r.Put("/{id}", h.UpdateUser)
		r.Delete("/{id}", h.DeleteUser)
		r.Get("/", h.ListUsers)
	})
}

type CreateUserRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (r *CreateUserRequest) Bind(_ *http.Request) error {
	return nil
}

// CreateUser godoc
// @Summary     Create a user
// @Description Create a user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user body CreateUserRequest true "User object"
// @Success     201 {object} model.User
// @Router     /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := render.Bind(r, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	user := &model.User{
		Login:    req.Login,
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	if err := h.service.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, user)
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body model.User true "User object"
// @Success 200 {object} model.User
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	user := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = id

	if err := h.service.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListUsers godoc
// @Summary List users
// @Description List users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, model.NewUserListResponse(users))
}
