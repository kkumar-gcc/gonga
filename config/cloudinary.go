package config

import (
	"os"
)

// CloudinaryConfig represents the configuration required to connect to the Cloudinary API.
type CloudinaryConfig struct {
	CloudName   string
	APIKey      string
	APISecret   string
	Environment string
}

func LoadCloudinaryConfig() *CloudinaryConfig {
	return &CloudinaryConfig{
		CloudName:   os.Getenv("CLOUDINARY_CLOUD_NAME"),
		APIKey:      os.Getenv("CLOUDINARY_API_KEY"),
		APISecret:   os.Getenv("CLOUDINARY_API_SECRET"),
		Environment: os.Getenv("API_ENV"),
	}
}
