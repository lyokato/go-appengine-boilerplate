package server

import (
	"app/config"
	"app/controller/admin"
	"app/controller/api"
	"app/controller/site"
	"app/middleware"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func New(cnf *config.Config) *gin.Engine {

	r := gin.New()

	r.Use(middleware.ConfigAccessor(cnf))
	r.Use(middleware.AppEngineContextAccessor())

	// TLS settins is now in app.yaml: using 'secure' parameter
	// r.Use(middleware.TLSOnly(cnf))

	admin.Setup(r.Group("/admin"), cnf)
	api.Setup(r.Group("/api"), cnf)
	site.Setup(r.Group("/site"), cnf)

	return r
}
