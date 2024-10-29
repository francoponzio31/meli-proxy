package app

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	AppHost       string
	AppPort       string
	MeliAPIHost   string
	RateLimiter   int
	TimeWindow    int
	RedisHost     string
	RedisPort     string
	RedisPassword string
}

var config *Config
var once sync.Once

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("Failed to load .env file")
		}

		config = &Config{
			AppHost:     getEnv("HOST", "0.0.0.0"),
			AppPort:     getEnv("PORT", "8080"),
			MeliAPIHost: getEnv("MELI_API_HOST", "https://api.mercadolibre.com/"),
			RateLimiter: getEnvAsInt("RATE_LIMIT", 1000),
			TimeWindow:  getEnvAsInt("TIME_WINDOW", 60),
			RedisHost:   getEnv("REDIS_HOST", "localhost"),
			RedisPort:   getEnv("REDIS_PORT", "6379"),
		}
	})

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Warning: Failed to parse %s as integer. Using default value %d\n", key, defaultValue)
		return defaultValue
	}
	return value
}
