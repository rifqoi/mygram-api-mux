package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		if os.Getenv("ENV") == "TEST" {
			err := godotenv.Load("../.env")
			if err != nil {
				log.Fatal("Error loading .env file masuk")
			}
		}
	}

	return os.Getenv(key)
}
