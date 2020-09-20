package config

import "os"

// Config struct
type Config struct {
	DB *DBConfig
}

// DBConfig struct
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Host     string
	Port     string
	Charset  string // default for mysql "utf8" not required for postgres
}

// GetConfig func
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  os.Getenv("DB_DRIVER"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Charset:  "utf8",
		},
	}
}
