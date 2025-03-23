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
		log.Fatal("[ERROR]: Error while reading data from .env$", err)
	}

	return &Config{
		PORT: os.Getenv("APP_PORT"),
	}
}
