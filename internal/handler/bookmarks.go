package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid/v5"
	"github.com/syntaqx/api/internal/model"
)

const (
	BookmarksURLPrefix = "/bookmarks"
)

type BookmarksHandler struct {
}

func NewBookmarksHandler() *BookmarksHandler {
	return &BookmarksHandler{}
}

func (h *BookmarksHandler) RegisterRoutes(r chi.Router) {
	r.Get(BookmarksURLPrefix, h.ListBookmarks)
}

func (h *BookmarksHandler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	bookmarks := []*model.Bookmark{
		{
			ID:   uuid.Must(uuid.NewV4()),
			Name: "Programming Resources",
			Type: model.BookmarkTypeFolder,
			Children: []*model.Bookmark{
				{
					ID:   uuid.Must(uuid.NewV4()),
					Name: "Go",
					Type: model.BookmarkTypeBookmark,
					URL:  "https://golang.org",
					Tags: []string{"go", "golang"},
				},
				{
					ID:   uuid.Must(uuid.NewV4()),
					Name: "Node.js",
					Type: model.BookmarkTypeBookmark,
					URL:  "https://nodejs.org",
					Tags: []string{"node", "nodejs"},
				},
			},
		},
	}

	render.RenderList(w, r, model.NewBookmarkListResponse(bookmarks))
}
