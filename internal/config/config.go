package config

import "os"

type Config struct {
	Port string
}

func LoadConfig() *Config {
	return &Config{
		Port: getEnv("PORT", ":4000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
