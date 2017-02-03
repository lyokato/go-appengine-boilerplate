package helper

import (
	"app/config"

	"golang.org/x/net/context"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Get Config Struct
func CONF(c *gin.Context) *config.Config {
	return c.MustGet("config").(*config.Config)
}

// Get AppEngine Context
func CTX(c *gin.Context) context.Context {
	return c.MustGet("aec").(context.Context)
}
