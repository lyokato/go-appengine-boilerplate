package admin

import (
	h "app/controller_helper"
	"strconv"

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

	admin_id, err := strconv.ParseInt(c.Param("admin_id"), 10, 64)
	if err != nil {
		log.Infof(h.CTX(c), "[admin] admin_id is not a integer")
		h.Render(c, "admin/error.html",
			pongo2.Context{"message": "invalid parameter"})
		return
	}

	admin, err := admins.FindById(h.GOON(c), admin_id)
	if err != nil {
		log.Infof(h.CTX(c), "[admin] not found for this admin_id")
		h.Render(c, "admin/error.html",
			pongo2.Context{"message": "not found"})
		return
	}

	h.Render(c, "admin/detail.html",
		pongo2.Context{"admin": admin})
}

func createAdmin(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] create_admin")
}

func updateAdmin(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] update_admin")

	admin_id, err := strconv.ParseInt(c.Param("admin_id"), 10, 64)
	if err != nil {
		log.Infof(h.CTX(c), "[admin] admin_id is not a integer")
		h.Render(c, "admin/error.html",
			pongo2.Context{"message": "invalid parameter"})
		return
	}

	admin, err := admins.FindById(h.GOON(c), admin_id)
	if err != nil {
		log.Infof(h.CTX(c), "[admin] not found for this admin_id")
		h.Render(c, "admin/error.html",
			pongo2.Context{"message": "not found"})
		return
	}

	h.Render(c, "admin/update.html",
		pongo2.Context{"admin": admin})
}

func deleteAdmin(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] delete_admin")
}
