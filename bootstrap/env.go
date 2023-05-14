package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		pterm.Fatal.Printf("Error loading .env file: %s", err)
	}
}
