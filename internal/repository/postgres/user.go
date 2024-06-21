package postgres

import (
	"github.com/gofrs/uuid/v5"

	"github.com/syntaqx/api/internal/model"
	"github.com/syntaqx/api/internal/repository"
)

type UserRepository struct {
	db repository.DB
}

// Assert userRepository implements UserRepository interface at comiple time.
var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db repository.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *model.User) error {
	user.ID = uuid.Must(uuid.NewV4())
	return repo.db.Create(user).Error
}

func (repo *UserRepository) GetUserByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(user *model.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepository) DeleteUser(id uuid.UUID) error {
	return repo.db.Delete(&model.User{}, "id = ?", id).Error
}

func (repo *UserRepository) ListUsers() ([]*model.User, error) {
	var users []*model.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
