package config

import (
	"fmt"
	"os"
)

type Config struct {
	ServerConfig
	DbConfig
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Sslmode  string
}

type ServerConfig struct {
	Env       string
	Port      string
	JwtSecret string
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func GetEnv() *Config {
	return &Config{
		ServerConfig: ServerConfig{
			Env:       getEnv("APP_ENV", "development"),
			Port:      getEnv("PORT", "8080"),
			JwtSecret: getEnv("JWT_SECRET", "default_secret_key"),
		},
		DbConfig: DbConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "postgres"),
			User:     getEnv("DB_USER", "postgres"),
			Password: os.Getenv("DB_PASSWORD"),
			Sslmode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
}

func (db *DbConfig) dsn() string {
	return fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=%s TimeZone=UTC",
		db.Host, db.User, db.Password, db.Name, db.Port, db.Sslmode)
}
