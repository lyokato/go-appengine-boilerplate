package middleware

import (
	"app/config"

	"github.com/mjibson/goon"

	"google.golang.org/appengine"
	gin "gopkg.in/gin-gonic/gin.v1"
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

func GoonAccessor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("goon", goon.NewGoon(c.Request))
		c.Next()
	}
}
