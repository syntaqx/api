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
	r.Get("/", h.GetRoot)
}

type RootResponse struct {
	Message string `json:"message"`
}

func (resp *RootResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// GetRoot godoc
// @Summary      Index
// @Description  get the index route
// @Tags         root
// @Accept       json
// @Produce      json
// @Success      200  {object}  RootResponse
// @Router       / [get]
func (h *RootHandler) GetRoot(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Render(w, r, &RootResponse{Message: "Hello, World!"})
}
