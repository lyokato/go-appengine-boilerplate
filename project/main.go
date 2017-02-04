package main

import (
	"net/http"

	"app/config"
	"app/middleware"
	"app/service/admin"
	"app/service/api"
	"app/service/site"

	"google.golang.org/appengine"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func setupEngine(cnf *config.Config) *gin.Engine {

	r := gin.New()

	r.Use(middleware.ConfigAccessor(cnf))
	r.Use(middleware.AppEngineContextAccessor())
	r.Use(middleware.GoonAccessor())

	// TLS settins is now in app.yaml: using 'secure' parameter
	// r.Use(middleware.TLSOnly(cnf))

	admin.Setup(r.Group("/admin"), cnf)
	api.Setup(r.Group("/api"), cnf)
	site.Setup(r.Group("/site"), cnf)

	return r
}

func main() {
	cnf := config.LoadFromFile("static/config/dev.toml")
	http.Handle("/", setupEngine(cnf))
	appengine.Main()
}
