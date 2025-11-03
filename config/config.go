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
	// 🧩 Load .env file without overwriting existing environment variables
	// (so host env always takes precedence)
	_ = godotenv.Load("config/.env")

	// If .env is missing, just log it (don’t panic)
	if _, err := os.Stat("config/.env"); err != nil {
		log.Println("📦 No local .env found — using system environment variables")
	} else {
		log.Println("✅ Loaded config/.env file")
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

	log.Printf("🔧 Loaded configuration: DB_HOST=%s, DB_NAME=%s, PORT=%s",
		cfg.DBHost, cfg.DBName, cfg.Port)

	return cfg
}

func getenv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}
