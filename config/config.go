package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func return env value
func Config(key string) string {
	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)
}