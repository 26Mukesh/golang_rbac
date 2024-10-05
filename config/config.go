package config

import (
	"log"

	"github.com/joho/godotenv"
)

// load environment variables file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}
