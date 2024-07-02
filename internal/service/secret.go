package service

import (
	"errors"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/syntaqx/api/internal/model"
	"github.com/syntaqx/api/internal/repository"
	"github.com/syntaqx/api/internal/util"
)

type SecretService interface {
	CreateSecret(secret string) (string, error)
	RetrieveSecret(id string) (string, error)
}

type secretService struct {
	repo repository.SecretRepository
}

func NewSecretService(repo repository.SecretRepository) SecretService {
	return &secretService{repo: repo}
}

func (s *secretService) CreateSecret(secret string) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	encryptedSecret, err := util.Encrypt(id.String(), secret)
	if err != nil {
		return "", err
	}

	newSecret := &model.Secret{
		ID:        id.String(),
		Secret:    encryptedSecret,
		Used:      false,
		CreatedAt: time.Now().Unix(),
	}

	if err := s.repo.Save(newSecret); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *secretService) RetrieveSecret(id string) (string, error) {
	storedSecret, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}

	if storedSecret.Used {
		return "", errors.New("secret already used")
	}

	decryptedSecret, err := util.Decrypt(id, storedSecret.Secret)
	if err != nil {
		return "", err
	}

	storedSecret.Used = true
	if err := s.repo.Update(storedSecret); err != nil {
		return "", err
	}

	return decryptedSecret, nil
}

// Compile-time assertion to ensure secretService implements SecretService
var _ SecretService = (*secretService)(nil)
