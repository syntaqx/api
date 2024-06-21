package repository

import (
	"github.com/gofrs/uuid/v5"

	"github.com/syntaqx/api/internal/model"
)

//go:generate go run github.com/matryer/moq -pkg mock -out ./mock/user_repository.go . UserRepository

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id uuid.UUID) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uuid.UUID) error
	ListUsers() ([]*model.User, error)
}
