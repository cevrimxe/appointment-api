package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	App      AppConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	URL      string
}

type ServerConfig struct {
	Port string
}

type JWTConfig struct {
	Secret string
}

type AppConfig struct {
	Environment string
	Debug       bool
}

func Load() *Config {
	// Load .env file if exists
	if err := godotenv.Load("config.env"); err != nil {
		log.Println("No config.env file found, using environment variables")
	}

	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "password"),
		Name:     getEnv("DB_NAME", "appointment_db"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Build database URL
	dbConfig.URL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.SSLMode)

	return &Config{
		Database: dbConfig,
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-secret-key"),
		},
		App: AppConfig{
			Environment: getEnv("APP_ENV", "development"),
			Debug:       getEnv("APP_DEBUG", "false") == "true",
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
