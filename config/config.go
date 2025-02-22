package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// Load configuration from env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Access environment variables
	dbHost := os.Getenv("DB_HOST")
	// Log or use these variables as needed
	log.Println("Database host:", dbHost)
}
