package controller_helper

import (
	"app/config"

	"github.com/mjibson/goon"

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

func GOON(c *gin.Context) goon.Goon {
	return c.MustGet("goon").(goon.Goon)
}
