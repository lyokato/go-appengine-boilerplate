package top_controller

import (
	"app/template"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"

	"google.golang.org/appengine/log"

	h "app/controller/helper"
)

func showIndex(c *gin.Context) {

	log.Infof(h.CTX(c), "top::/")

	err := template.Render(c.Writer, "index.html",
		pongo2.Context{"title": h.CONF(c).Web.Address()})

	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}

}
