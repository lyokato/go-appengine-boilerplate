package admin

import (
	h "app/controller_helper"
	"app/model/admins"
	"fmt"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-contrib/sessions"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func showIndexPage(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] show_index")

	msg := fmt.Sprintf("Hello, %s", ADMIN(c).Username)

	h.Render(c, "admin/index.html",
		pongo2.Context{"title": msg})
}

func signIn(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] sign_in")

	_, exists := c.Get(StashKeyAdmin)
	if exists {
		log.Infof(h.CTX(c), "[admin] already signed in, redirect to top page")
		c.Redirect(http.StatusFound, "/admin/")
		return
	}

	url, err := user.LoginURL(h.CTX(c), "/admin/sign_in_cb")
	if err != nil {
		log.Warningf(h.CTX(c), "[admin] failed to obtain login URL")

		h.Render(c, "admin/error.html",
			pongo2.Context{})
		return
	}

	c.Redirect(http.StatusFound, url)
}

func signInCallback(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] sign_in_callback")

	_, exists := c.Get(StashKeyAdmin)
	if exists {
		log.Infof(h.CTX(c), "[admin] already signed in, redirect to top page")
		c.Redirect(http.StatusFound, "/admin/")
		return
	}

	u := user.Current(h.CTX(c))

	log.Infof(h.CTX(c), "[admin] google-account '%s' found", u.Email)

	admin, err := admins.FindByEmail(h.GOON(c), u.Email)
	if err == nil {

		log.Infof(h.CTX(c), "[admin] admin found, redirect to top page")

		sess := sessions.Default(c)
		sess.Set(SessionKeyLoginAdminId, admin.AdminId)
		sess.Save()

		c.Redirect(http.StatusFound, "/admin/")
		return
	}

	if h.CONF(c).Admin.InitialAccount != u.Email {
		log.Infof(h.CTX(c), "[admin] this google-account is not set as admin")

		// TODO showError
		c.Redirect(http.StatusFound, "/admin/sign_in_cb")
		return
	}

	admin, err = admins.Create(h.GOON(c), u.ID, u.Email)
	if err != nil {
		log.Warningf(h.CTX(c), "[admin] failed to create new admin")

		// TODO showError
		c.Redirect(http.StatusFound, "/admin/sign_in")
		return
	}

	sess := sessions.Default(c)
	sess.Set(SessionKeyLoginAdminId, admin.AdminId)
	sess.Save()

	c.Redirect(http.StatusFound, "/admin/")
}

func signOut(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] sign_out")

	user.LogoutURL(h.CTX(c), "/admin/sign_out_cb")
}

func signOutCallback(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] sign_out_callback")

	sess := sessions.Default(c)
	sess.Delete(SessionKeyLoginAdminId)
	sess.Save()

	c.Redirect(http.StatusFound, "/admin/sign_in")
}
