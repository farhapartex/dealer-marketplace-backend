package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Settings struct {
	EncryptKey      string
	AuthCodeExpiry  string
	JWTSecretKey    string
	Database        DatabaseSettings
}

type DatabaseSettings struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
}

var AppSettings *Settings

func LoadSettings() error {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return fmt.Errorf("invalid DB_PORT: %w", err)
	}

	AppSettings = &Settings{
		EncryptKey:     getEnv("ENCRYPT_KEY", ""),
		AuthCodeExpiry: getEnv("AUTH_CODE_EXPIRY", "1HR"),
		JWTSecretKey:   getEnv("JWT_SECRET_KEY", ""),
		Database: DatabaseSettings{
			Name:     getEnv("DB_NAME", "dealer_db"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     port,
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
		},
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (s *Settings) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Database.Host,
		s.Database.Port,
		s.Database.User,
		s.Database.Password,
		s.Database.Name,
	)
}
