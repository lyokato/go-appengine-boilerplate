package template

import (
	"io"

	"github.com/flosch/pongo2"
)

func Render(w io.Writer, name string, data interface{}) error {
	tpl, err := pongo2.FromCache(filePath(name))
	if err != nil {
		return err
	}
	ctx := data.(pongo2.Context)
	err = tpl.ExecuteWriter(ctx, w)
	if err != nil {
		return err
	}
	return nil
}

func filePath(name string) string {
	return "static/templates/" + name
}
