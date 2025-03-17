package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bharabhi01/shorturl-go/api/handlers"
	"github.com/bharabhi01/shorturl-go/api/middleware"
	"github.com/bharabhi01/shorturl-go/config"
	"github.com/bharabhi01/shorturl-go/internal/repository"
	"github.com/bharabhi01/shorturl-go/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	db, err := repository.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Connect to Redis
	cache, err := repository.NewCache(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Create repository
	urlRepo := repository.NewURLRepository(db, cache)

	// Create service
	urlService := service.NewURLService(urlRepo, cfg.BaseURL)

	// Create handler
	urlHandler := handlers.NewURLHandler(urlService)

	// Create router
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())

	apiGroup := router.Group("/api")
	apiGroup.Use(middleware.RateLimiter(cache.Client, cfg.RateLimit, time.Minute))
	apiGroup.POST("/urls", urlHandler.CreateShortURL)

	router.GET("/:shortCode", urlHandler.RedirectToLongURL)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
