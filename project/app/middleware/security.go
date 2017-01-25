package middleware

import (
	"app/config"

	"github.com/gin-gonic/gin"
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
