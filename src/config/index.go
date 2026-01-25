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
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	PORT = os.Getenv("PORT")
	MONGODB_URI = os.Getenv("MONGODB_URI")
}
