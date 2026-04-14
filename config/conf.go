package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	Port   string
}

func LoadEnv() Config {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}
	return Config{
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		Port:   os.Getenv("PORT"),
	}
}