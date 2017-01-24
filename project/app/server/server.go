package server

import (
	"app/config"
	"app/controller/top_controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(cnf *config.Config) *gin.Engine {

	r := gin.New()

	r.Use(configAccessor(cnf))
	r.Use(appEngineContextAccessor())

	top_controller.Route(r.Group("/top"))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, cnf.Web.Address())
	})

	return r
}
