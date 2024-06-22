package model

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func NewUserListResponse(users []*User) []render.Renderer {
	list := []render.Renderer{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}
