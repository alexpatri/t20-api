package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config armazena as configurações da aplicação
type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

// ServerConfig armazena as configurações do servidor HTTP
type ServerConfig struct {
	Port string
}

// DatabaseConfig armazena as configurações do banco de dados
type DatabaseConfig struct {
	URI          string
	DatabaseName string
}

// LoadConfig carrega as configurações do ambiente
func LoadConfig(env string) *Config {
	if env == "file" {
		godotenv.Load()
	}

	return &Config{
		Server: &ServerConfig{
			Port: getEnv("SERVER_PORT", "8000"),
		},
		Database: &DatabaseConfig{
			URI:          getEnv("MONGODB_URI", "mongodb://localhost:27017"),
			DatabaseName: getEnv("MONGODB_DATABASE", ""),
		},
	}
}

// getEnv obtém uma variável de ambiente ou retorna um valor padrão
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" {
		return defaultValue
	}
	return value
}
