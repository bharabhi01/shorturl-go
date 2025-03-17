package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RateLimiter(redisClient *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rate_limit:" + ip

		ctx := context.Background()
		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			log.Printf("Redis error: %v", err)
			c.Next()
			return
		}

		if count == 1 {
			if err := redisClient.Expire(ctx, key, window).Err(); err != nil {
				log.Printf("Failed to set expiry for rate limit key: %v", err)
			}
		}

		if count > int64(limit) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
