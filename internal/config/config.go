package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort int
	LogLevel int
	AppEnv   string
}

var once sync.Once

var config *Config

func Get() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using environment variable")
		}

		config = &Config{
			HttpPort: getEnvAsInt("HTTP_PORT", 8080),
			LogLevel: getEnvAsInt("LOG_LEVEL", 1),
			AppEnv:   getEnv("APP_ENV", "development"),
		}
	})

	return config
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Printf("Could not find %s Key, thus defaulting to %s\n", key, defaultValue)
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	log.Printf("Could not parse value for Key %s, thus defaulting to %d\n", key, defaultValue)
	return defaultValue
}
