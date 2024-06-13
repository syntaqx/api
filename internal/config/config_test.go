package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfigDefaults(t *testing.T) {
	cfg := NewConfig()
	assert.Equal(t, cfg.Port, "8080")
}
