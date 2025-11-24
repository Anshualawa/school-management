package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBUser     string
	DBPass     string
	DBName     string
	DBPort     string
	JWTSecrete string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", "Alawa@3282"),
		DBName:     getEnv("DB_NAME", "my_school"),
		DBPort:     getEnv("DB_PORT", "3306"),
		JWTSecrete: getEnv("JWT_SECRET", "alawasuperscretedev"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
