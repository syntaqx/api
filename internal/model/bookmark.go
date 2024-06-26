package model

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/gofrs/uuid/v5"
)

type BookmarkType string

const (
	BookmarkTypeFolder   BookmarkType = "folder"
	BookmarkTypeBookmark BookmarkType = "bookmark"
)

type Bookmark struct {
	ID        uuid.UUID    `gorm:"primaryKey" json:"id"`
	Name      string       `json:"name"`
	Type      BookmarkType `json:"type"`
	URL       string       `json:"url,omitempty"`
	Tags      []string     `json:"tags"`
	Children  []*Bookmark  `json:"children,omitempty"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func (b *Bookmark) Bind(r *http.Request) error {
	return nil
}

func (b *Bookmark) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func NewBookmarkListResponse(bookmarks []*Bookmark) []render.Renderer {
	list := []render.Renderer{}
	for _, bookmark := range bookmarks {
		list = append(list, bookmark)
	}
	return list
}
