package model

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/gofrs/uuid/v5"
)

type Game struct {
	ID   uuid.UUID `gorm:"primaryKey" json:"id"`
	Name string    `json:"name"`
}

func (g *Game) Bind(r *http.Request) error {
	return nil
}

func (g *Game) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func NewGameListResponse(games []*Game) []render.Renderer {
	list := []render.Renderer{}
	for _, game := range games {
		list = append(list, game)
	}
	return list
}
