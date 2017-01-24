package main

import (
	"app/config"
	"app/server"

	"net/http"

	"google.golang.org/appengine"
)

func main() {
	cnf := config.LoadFromFile("static/config/dev.toml")
	http.Handle("/", server.New(cnf))
	appengine.Main()
}
