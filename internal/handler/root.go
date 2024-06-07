package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RootHandler struct {
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.Root)
}

func (h *RootHandler) Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World!")
}
