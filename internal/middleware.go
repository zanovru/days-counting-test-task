package internal

import "github.com/gin-gonic/gin"

// PingPongCheckerMiddleware checks X-PING key and ping value in HTTP header
func PingPongCheckerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("X-PING")
		if value == "ping" {
			c.Header("X-PONG", "pong")
		}
		c.Next()
	}
}
