package helper

import (
	"app/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// Get Config Struct
func CONF(c *gin.Context) *config.Config {
	return c.MustGet("config").(*config.Config)
}

// Get AppEngine Context
func CTX(c *gin.Context) context.Context {
	return c.MustGet("aec").(context.Context)
}
