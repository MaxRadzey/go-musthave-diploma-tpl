package config

import (
	"flag"
	"os"
)

type Config struct {
	RunAddress           string
	DatabaseDSN          string
	AccrualSystemAddress string
	LogLevel             string
	CookieSecret         string
}

func New() *Config {
	return &Config{
		RunAddress:           "localhost:8080",
		LogLevel:             "info",
		DatabaseDSN:          "postgres://shortener:shortener@localhost:5432/shortener",
		CookieSecret:         "dev-secret",
		AccrualSystemAddress: "",
	}
}

func ParseEnv(config *Config) {
	if RunAddress := os.Getenv("SERVER_ADDRESS"); RunAddress != "" {
		config.RunAddress = RunAddress
	}
	if LogLevel := os.Getenv("LOG_LEVEL"); LogLevel != "" {
		config.LogLevel = LogLevel
	}
	if DatabaseDSN := os.Getenv("DATABASE_DSN"); DatabaseDSN != "" {
		config.DatabaseDSN = DatabaseDSN
	}
	if s := os.Getenv("COOKIE_SECRET"); s != "" {
		config.CookieSecret = s
	}
	if s := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); s != "" {
		config.AccrualSystemAddress = s
	}
}

func ParseFlags(config *Config) {
	flag.StringVar(&config.RunAddress, "a", config.RunAddress, "address and port to run server")
	flag.StringVar(&config.LogLevel, "l", config.LogLevel, "log level")
	flag.StringVar(&config.DatabaseDSN, "d", config.DatabaseDSN, "database connection string")
	flag.StringVar(&config.AccrualSystemAddress, "r", config.AccrualSystemAddress, "accrual system address")

	flag.Parse()
}
