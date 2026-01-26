package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT        string
	MONGODB_URI string
)

func LoadConfig() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
	} else {
		log.Println(".env file not found, using environment variables")
	}

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	MONGODB_URI = os.Getenv("MONGODB_URI")
}
