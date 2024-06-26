package model

import (
	"testing"

	"github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkBind(t *testing.T) {
	bookmark := &Bookmark{}
	err := bookmark.Bind(nil)
	assert.Nil(t, err)
}

func TestBookmarkRender(t *testing.T) {
	bookmark := &Bookmark{}
	err := bookmark.Render(nil, nil)
	assert.Nil(t, err)
}

func TestNewBookmarkListResponse(t *testing.T) {
	bookmarks := []*Bookmark{
		{
			ID:   uuid.Nil,
			Name: "test",
		},
	}

	list := NewBookmarkListResponse(bookmarks)
	assert.Equal(t, 1, len(list))
}
