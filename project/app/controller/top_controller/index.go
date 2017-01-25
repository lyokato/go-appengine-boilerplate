package top_controller

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"

	"google.golang.org/appengine/log"

	h "app/controller/helper"
)

func showIndex(c *gin.Context) {

	log.Infof(h.CTX(c), "top::/")

	h.Render(c, "index.html",
		pongo2.Context{"title": h.CONF(c).Web.Address()})

}
