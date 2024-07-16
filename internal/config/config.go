package config

import "github.com/syntaqx/env"

type Config struct {
	FQDN           string `env:"FQDN,default=localhost"`
	HTTPS          bool   `env:"HTTPS,default=false"`
	Port           string `env:"PORT,default=8080"`
	DatabaseURL    string `env:"DATABASE_URL,default=postgres://postgres:postgres@localhost:5432/api?sslmode=disable"`
	WeatherAPIHost string `env:"WEATHER_API_HOST,default=https://api.weatherapi.com"`
	WeatherAPIKey  string `env:"WEATHER_API_KEY"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
