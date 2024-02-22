package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	DSN string
}

func validateDatabaseConfig() error {
	if os.Getenv("DB_DSN") == "" {
		return fmt.Errorf("DB_DSN config is not set")
	}
	return nil
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	err := validateDatabaseConfig()

	if err != nil {
		return nil, err
	}

	return &DatabaseConfig{
		DSN: os.Getenv("DB_DSN"),
	}, err
}
