package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	RedisURL    string
	ServerPort  string
	BaseURL     string
	RateLimit   int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found, using default env variables")
	}

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "urlshortner")

	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisPassword := getEnv("REDIS_PASSWORD", "")

	serverPort := getEnv("SERVER_PORT", "8080")
	baseURL := getEnv("BASE_URL", "http://localhost:8080")

	rateLimit, err := strconv.Atoi(getEnv("RATE_LIMIT", "100"))
	if err != nil {
		return nil, fmt.Errorf("invalid rate limit: %v", err)
	}

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	redisURL := fmt.Sprintf("redis://%s:%s@%s:%s/0", "", redisPassword, redisHost, redisPort)

	return &Config{
		DatabaseURL: dbURL,
		RedisURL:    redisURL,
		ServerPort:  serverPort,
		BaseURL:     baseURL,
		RateLimit:   rateLimit,
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
