package main

import (
	"config"
	"net/http"
	"server"

	"google.golang.org/appengine"
)

func main() {
	cnf := config.LoadFromFile("static/config/dev.toml")
	http.Handle("/", server.New(cnf))
	appengine.Main()
}
