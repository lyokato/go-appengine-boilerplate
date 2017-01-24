package top_controller

import (
	"app/template"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	err := template.Render(c.Writer, "index.html",
		pongo2.Context{"title": "hogehoge"})
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
}
