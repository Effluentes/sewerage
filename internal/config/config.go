package config

import (
	"fmt"
	"os"
)

type Config struct {
    PostgresURL    string `env:"POSTGRES_URL" required:"true"`
    TimescaleURL   string `env:"TIMESCALE_URL" required:"true"`
    RedisAddr      string `env:"REDIS_ADDR" required:"true"`
    RedisPassword  string `env:"REDIS_PASSWORD"`
    RedisDB        int    `env:"REDIS_DB" default:"0"`
}

func getEnv(key string) string {
	// Get the value of the key from the environment
	value := os.Getenv(key)
	if value == "" {
		// If the value is empty, return an error
		return fmt.Errorf("environment variable %s is not set", key)
	}
	return value
}

