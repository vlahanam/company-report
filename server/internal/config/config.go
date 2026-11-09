package config

import "os"

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

func GetEnv() *Config {
	return &Config{
		ServerConfig{
			Env:       os.Getenv("APP_ENV"),
			Port:      os.Getenv("PORT"),
			JwtSecret: os.Getenv("JWT_SECRET"),
		},
		DbConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Sslmode:  os.Getenv("DB_SSLMODE"),
		},
	}
}
