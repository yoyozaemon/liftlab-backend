package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT        string
	MONGODB_URI string
	JWT_SECRET  string
	DB_NAME     string
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
	if MONGODB_URI == "" {
		log.Panic("FATAL: MONGODB_URI is not set in environment")
	}

	JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Panic("FATAL: JWT_SECRET is not set in environment")
	}

	DB_NAME = os.Getenv("DB_NAME")
	if DB_NAME == "" {
		DB_NAME = "liftlab"
	}
}
