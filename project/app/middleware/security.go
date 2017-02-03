package middleware

import (
	"app/config"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func TLSOnly(cnf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.TLS != nil {
			c.Next()
		} else {
			c.AbortWithStatus(400)
		}
	}
}

func DefaultSecureHeaders(cnf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "deny")
		c.Next()
	}
}
