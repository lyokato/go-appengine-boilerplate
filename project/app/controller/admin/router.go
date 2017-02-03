package admin

import (
	"app/config"
	h "app/controller/helper"
	"app/middleware"

	"github.com/flosch/pongo2"
	"google.golang.org/appengine/log"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func Setup(r *gin.RouterGroup, cnf *config.Config) {

	r.Use(middleware.Sessions(cnf.AdminSession))

	r.GET("/", showIndexPage)
	r.GET("/login", showLoginPage)
	r.POST("/login", login)
	r.GET("/logout", logout)

}

func showIndexPage(c *gin.Context) {
	log.Infof(h.CTX(c), "api::top::showIndex")

	h.Render(c, "api/index.html",
		pongo2.Context{"title": h.CONF(c).Web.Address()})
}

func showLoginPage(c *gin.Context) {
}

func login(c *gin.Context) {

}

func logout(c *gin.Context) {

}
