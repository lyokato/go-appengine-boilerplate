package server

import (
	"config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(cnf *config.Config) *gin.Engine {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, cnf.Web.Address())
	})
	return r
}
