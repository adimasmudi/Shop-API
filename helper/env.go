package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvDBURL() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_URL")
}