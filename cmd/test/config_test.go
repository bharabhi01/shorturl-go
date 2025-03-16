package main

import (
	"fmt"
	"github.com/bharabhi01/shorturl-go/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config Loaded: %+v\n", cfg)
	fmt.Printf("Database URL: %s\n", cfg.DatabaseURL)
	fmt.Printf("Redis URL: %s\n", cfg.RedisURL)
	fmt.Printf("Server Port: %s\n", cfg.ServerPort)
	fmt.Printf("Base URL: %s\n", cfg.BaseURL)
}