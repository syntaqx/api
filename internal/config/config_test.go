package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfigDefaults(t *testing.T) {
	cfg, err := NewConfig()
	assert.NoError(t, err)
	assert.Equal(t, cfg.Port, "8080")
}
