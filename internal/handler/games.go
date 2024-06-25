package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid/v5"

	"github.com/syntaqx/api/internal/model"
)

const (
	GamesURLPrefix = "/games"
)

type GamesHandler struct {
}

func NewGamesHandler() *GamesHandler {
	return &GamesHandler{}
}

func (h *GamesHandler) RegisterRoutes(r chi.Router) {
	r.Get(GamesURLPrefix, h.ListGames)
}

func (h *GamesHandler) ListGames(w http.ResponseWriter, r *http.Request) {
	games := []*model.Game{
		{
			ID:   uuid.Must(uuid.NewV4()),
			Name: "World of Warcraft",
		},
		{
			ID:   uuid.Must(uuid.NewV4()),
			Name: "Apex Legends",
		},
	}

	render.RenderList(w, r, model.NewGameListResponse(games))
}
