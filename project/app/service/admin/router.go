package admin

import (
	"app/config"
	"app/middleware"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func Setup(r *gin.RouterGroup, cnf *config.Config) {

	r.Use(middleware.Sessions(cnf.AdminSession))
	r.Use(authenticator("/admin/sign_in", []string{"/admin/sign_in", "/admin/sign_in_cb"}))

	r.GET("/", showIndexPage)
	r.GET("/sign_in", signIn)
	r.GET("/sign_in_cb", signInCallback)
	r.GET("/sign_out", signOut)
	r.GET("/sign_out_cb", signOutCallback)
}
