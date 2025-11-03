package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	JWTSecret   string
	DatabaseURL string // constructed dynamically
}

func Load() *Config {
	// Load .env file
	if err := godotenv.Load("config/.env"); err != nil {
		log.Println("⚠️ No .env file found, using environment variables")
	}

	cfg := &Config{
		Port:       getenv("PORT", "8080"),
		DBHost:     getenv("DB_HOST", "localhost"),
		DBPort:     getenv("DB_PORT", "5432"),
		DBUser:     getenv("DB_USER", "postgres"),
		DBPassword: getenv("DB_PASSWORD", ""),
		DBName:     getenv("DB_NAME", "postgres"),
		JWTSecret:  getenv("JWT_SECRET", "supersecret"),
	}

	cfg.DatabaseURL = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	log.Printf("Loaded configuration: DB_HOST=%s, DB_NAME=%s", cfg.DBHost, cfg.DBName)
	return cfg
}

func getenv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
