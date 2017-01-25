package middleware

import (
	"app/config"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

func ConfigAccessor(cnf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", cnf)
		c.Next()
	}
}

func AppEngineContextAccessor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("aec", appengine.NewContext(c.Request))
		c.Next()
	}
}
