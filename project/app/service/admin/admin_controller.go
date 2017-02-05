package admin

import (
	h "app/controller_helper"

	"app/model/admins"

	"github.com/flosch/pongo2"
	"google.golang.org/appengine/log"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func routeAdminActions(r *gin.RouterGroup) {
	sub := r.Group("admin")
	sub.GET("/", showAdminList)
	sub.GET("/new", createAdmin)
	sub.GET("/detail/:admin_id", showAdminDetail)
	sub.POST("/detail/:admin_id", updateAdmin)
}

func showAdminList(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] show_admin_list")

	list, err := admins.All(h.CTX(c))
	if err != nil {
		h.Render(c, "admin/error.html",
			pongo2.Context{})
		return
	}

	h.Render(c, "admin/admins.html",
		pongo2.Context{"admins": list})
}

func showAdminDetail(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] show_admin_detail")
}

func createAdmin(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] create_admin")
}

func updateAdmin(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] update_admin")
}

func deleteAdmin(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] delete_admin")
}
