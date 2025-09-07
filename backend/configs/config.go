package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("[WARNING]: .env file not found, using environment variables")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = ":8080"
		}
	}

	return &Config{
		PORT: port,
	}
}
