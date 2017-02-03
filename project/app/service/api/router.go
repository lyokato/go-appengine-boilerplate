package api

import (
	"app/config"
	h "app/controller_helper"

	"net/http"

	"google.golang.org/appengine/log"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func Setup(r *gin.RouterGroup, cnf *config.Config) {
	r.GET("/", showIndex)
}

func showIndex(c *gin.Context) {

	log.Infof(h.CTX(c), "api::top::showIndex")

	c.JSON(http.StatusOK, gin.H{"message": "hello, world"})
}
