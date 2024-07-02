package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/syntaqx/api/internal/service"
)

const SecretURLPrefix = "/secrets"

type SecretHandler struct {
	service service.SecretService
}

func NewSecretHandler(service service.SecretService) *SecretHandler {
	return &SecretHandler{service: service}
}

func (h *SecretHandler) RegisterRoutes(r chi.Router) {
	r.Post(SecretURLPrefix, h.CreateSecret)
	r.Get(SecretURLPrefix+"/{secretId}", h.RetrieveSecret)
}

type CreateSecretRequest struct {
	Secret string `json:"secret"`
}

// @Summary      Create a secret
// @Description  Create a secret
// @Router       /secrets [post]
// @Tags         secrets
// @Accept       json
// @Produce      json
// @Param        secret  body  CreateSecretRequest  true  "Secret"
// @Success      200  {object}  map[string]string
func (h *SecretHandler) CreateSecret(w http.ResponseWriter, r *http.Request) {
	var req CreateSecretRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateSecret(req.Secret)
	if err != nil {
		http.Error(w, "failed to create secret", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{"id": id})
}

// @Summary      Retrieve a secret
// @Description  Retrieve a secret
// @Router       /secrets/{secretId} [get]
// @Tags         secrets
// @Accept       json
// @Produce      json
// @Param        secretId  path  string  true  "Secret ID"
// @Success      200  {object}  map[string]string
func (h *SecretHandler) RetrieveSecret(w http.ResponseWriter, r *http.Request) {
	secretId := chi.URLParam(r, "secretId")

	secret, err := h.service.RetrieveSecret(secretId)
	if err != nil {
		http.Error(w, "failed to retrieve secret", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{"secret": secret})
}
