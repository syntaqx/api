package config

import "github.com/syntaqx/env"

type Config struct {
	Host           string
	Port           string
	WeatherAPIHost string
	WeatherAPIKey  string
}

func NewConfig() *Config {
	return &Config{
		Host:           env.GetWithFallback("HOST", "localhost"),
		Port:           env.GetWithFallback("PORT", "8080"),
		WeatherAPIHost: env.GetWithFallback("WEATHER_API_HOST", "https://api.weatherapi.com"),
		WeatherAPIKey:  env.Get("WEATHER_API_KEY"),
	}
}
