package middleware

import (
	"github.com/gin-gonic/gin"
	"project_06/limiter"
)

func Limiter(l *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !l.Allow() {
			c.JSON(429, gin.H{
				"message": "too many requests",
				"status":  429,
			})
			c.Abort()
		}
		c.Next()
	}

}
