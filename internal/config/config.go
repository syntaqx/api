package config

import "github.com/syntaqx/env"

type Config struct {
	Port          string
	WeatherAPIKey string
}

func NewConfig() *Config {
	return &Config{
		Port:          env.GetWithFallback("PORT", "8080"),
		WeatherAPIKey: env.Get("WEATHER_API_KEY"),
	}
}
