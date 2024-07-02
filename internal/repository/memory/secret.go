package memory

import (
	"errors"
	"sync"

	"github.com/syntaqx/api/internal/model"
)

type SecretRepository struct {
	secrets map[string]*model.Secret
	mu      sync.Mutex
}

func NewSecretRepository() *SecretRepository {
	return &SecretRepository{
		secrets: make(map[string]*model.Secret),
	}
}

func (r *SecretRepository) Save(secret *model.Secret) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.secrets[secret.ID]; exists {
		return errors.New("secret already exists")
	}

	r.secrets[secret.ID] = secret
	return nil
}

func (r *SecretRepository) FindByID(id string) (*model.Secret, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	secret, exists := r.secrets[id]
	if !exists {
		return nil, errors.New("secret not found")
	}

	return secret, nil
}

func (r *SecretRepository) Update(secret *model.Secret) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.secrets[secret.ID]; !exists {
		return errors.New("secret not found")
	}

	r.secrets[secret.ID] = secret
	return nil
}
