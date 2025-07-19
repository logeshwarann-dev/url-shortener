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

func SetEnv(variable string, value string) {
	os.Setenv(variable, value)
}

func GetEnv(envName string) string {
	return os.Getenv(envName)
}
