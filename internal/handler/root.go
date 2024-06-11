package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type RootHandler struct {
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.Root)
}

type RootResponse struct {
	Message string `json:"message"`
}

func (resp *RootResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *RootHandler) Root(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Render(w, r, &RootResponse{Message: "Hello, World!"})
}
