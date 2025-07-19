package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(filepath string) {
	err := godotenv.Load(filepath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(envName string) string {
	return os.Getenv(envName)
}
