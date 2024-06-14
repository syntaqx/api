package model

import (
	"time"

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
