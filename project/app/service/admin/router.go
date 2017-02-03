package admin

import (
	"app/config"
	"app/middleware"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func Setup(r *gin.RouterGroup, cnf *config.Config) {

	r.Use(middleware.Sessions(cnf.AdminSession))
	r.Use(authenticator("/admin/signin", []string{"/admin/signin"}))

	r.GET("/", showIndexPage)
	r.GET("/signin", showSignInPage)
	r.POST("/signin", signIn)
	r.GET("/signout", signOut)
}
