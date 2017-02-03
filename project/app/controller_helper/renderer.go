package controller_helper

import (
	"app/template"

	"github.com/flosch/pongo2"
	gin "gopkg.in/gin-gonic/gin.v1"

	"google.golang.org/appengine/log"
)

func Render(c *gin.Context, file string,
	params pongo2.Context) error {

	err := template.Render(c.Writer, file, params)

	if err != nil {
		log.Errorf(CTX(c),
			"failed to render template", err.Error())

		c.String(500, "Internal Server Error")
		return err
	}

	return nil
}
