package service

import (
	"github.com/gofrs/uuid/v5"

	"github.com/syntaqx/api/internal/model"
	"github.com/syntaqx/api/internal/repository"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id uuid.UUID) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uuid.UUID) error
	ListUsers() ([]*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *model.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id uuid.UUID) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}

func (s *userService) ListUsers() ([]*model.User, error) {
	return s.repo.ListUsers()
}
