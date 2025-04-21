package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func InitialEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return nil
}