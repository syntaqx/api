package repository

import "github.com/syntaqx/api/internal/model"

type SecretRepository interface {
	Save(secret *model.Secret) error
	FindByID(id string) (*model.Secret, error)
	Update(secret *model.Secret) error
}
