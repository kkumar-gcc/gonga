package config

import (
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database     string
	Charset  string
}

func LoadDatabaseConfig() (*DatabaseConfig) {
	return &DatabaseConfig{
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database:      os.Getenv("DB_DATABASE"),
		Charset:  "utf8mb4",
	}
}


