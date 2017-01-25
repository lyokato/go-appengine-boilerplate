package server

import (
	"app/config"
	"app/controller/top_controller"
	"app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(cnf *config.Config) *gin.Engine {

	r := gin.New()

	r.Use(middleware.ConfigAccessor(cnf))
	r.Use(middleware.AppEngineContextAccessor())

	// TLS settins is now in app.yaml: using 'secure' parameter
	// r.Use(middleware.TLSOnly(cnf))

	top_controller.Route(r.Group("/top"))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, cnf.Web.Address())
	})

	return r
}
