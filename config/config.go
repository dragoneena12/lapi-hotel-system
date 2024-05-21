package config

import (
	"os"
)

type Config struct {
	Debug     bool
	DbName    string
	SQLDriver string
	Port      string
	JWTSecret string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	if os.Getenv("SQL_DRIVER") == "" {
		os.Setenv("SQL_DRIVER", "sqlite3")
	}
	if os.Getenv("DB_NAME") == "" {
		os.Setenv("DB_NAME", "development.db")
	}
	return &Config{
		Debug:     os.Getenv("DEBUG") == "true",
		SQLDriver: os.Getenv("SQL_DRIVER"),
		DbName:    os.Getenv("DB_NAME"),
		Port:      port,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
