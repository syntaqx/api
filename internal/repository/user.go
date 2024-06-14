package repository

import (
	"github.com/gofrs/uuid/v5"

	"github.com/syntaqx/api/internal/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id uuid.UUID) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uuid.UUID) error
	ListUsers() ([]*model.User, error)
}
