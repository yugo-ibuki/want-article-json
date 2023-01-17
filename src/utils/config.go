package config

import (
	"github.com/joho/godotenv"
	"os"
)

func GetApiKey() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	apiKey := os.Getenv("API_KEY")

	return apiKey, nil
}
