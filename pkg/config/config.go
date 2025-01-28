package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig  DBConfig
	JWTSecret string
}

type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     int
}

var (
	config *Config
)

func Get() *Config {
	if config == nil {
		cfg, err := Load()
		if err != nil {
			panic(fmt.Sprintf("Failed to load config: %v", err))
		}
		config = cfg
	}
	return config
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	cfg := &Config{
		DBConfig: DBConfig{
			Host:     getEnv("DB_HOST"),
			Name:     getEnv("DB_NAME"),
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			Port:     dbPort,
		},
		JWTSecret: getEnv("JWT_SECRET"),
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.DBConfig.Name == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	if c.DBConfig.User == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if c.DBConfig.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}
	if c.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}
	return nil
}

func getEnv(key string, fallback ...string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return ""
}
