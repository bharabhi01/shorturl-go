package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		gin.DefaultWriter.Write([]byte(
			time.Now().Format("27/02/2006 - 15:04:05") + " | " +
				method + " | " +
				path + " | " +
				string(rune(statusCode)) + " | " +
				latency.String() + "\n"))
	}
}
