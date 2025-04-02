package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port  string
	Token string
}

type Config struct {
	Server *ServerConfig
}

func loadConfig(path string) {
	fileName := ".env"

	if path != "" {
		fileName = path
	}

	if err := godotenv.Load(fileName); err != nil {
		log.Printf("No .env file found, using system environment variables. err: %s", err.Error())
	}
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func NewConfig(path string) *Config {
	loadConfig(path)

	defaultConfig := &Config{
		Server: &ServerConfig{
			Port:  GetEnv("PORT", "8080"),
			Token: GetEnv("TOKEN", ""),
		},
	}

	return &Config{
		Server: defaultConfig.Server,
	}
}
