package site

import (
	"app/config"
	h "app/controller_helper"
	"app/middleware"

	"github.com/flosch/pongo2"
	"google.golang.org/appengine/log"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func Setup(r *gin.RouterGroup, cnf *config.Config) {

	r.Use(middleware.Sessions(cnf.SiteSession))

	r.GET("/", showIndexPage)
}

func showIndexPage(c *gin.Context) {

	log.Infof(h.CTX(c), "site::top::showIndexPage")

	h.Render(c, "site/index.html",
		pongo2.Context{"title": "Hello, World"})

}
