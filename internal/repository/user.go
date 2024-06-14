package repository

import (
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"

	"github.com/syntaqx/api/internal/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id uuid.UUID) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uuid.UUID) error
	ListUsers() ([]*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user *model.User) error {
	user.ID = uuid.Must(uuid.NewV4())
	return repo.db.Create(user).Error
}

func (repo *userRepository) GetUserByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) UpdateUser(user *model.User) error {
	return repo.db.Save(user).Error
}

func (repo *userRepository) DeleteUser(id uuid.UUID) error {
	return repo.db.Delete(&model.User{}, "id = ?", id).Error
}

func (repo *userRepository) ListUsers() ([]*model.User, error) {
	var users []*model.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
