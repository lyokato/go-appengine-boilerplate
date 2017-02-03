package admin

import (
	h "app/controller_helper"
	"app/model/admins"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/flosch/pongo2"
	"github.com/gin-contrib/sessions"
	"google.golang.org/appengine/log"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func showIndexPage(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] show_index")

	h.Render(c, "admin/index.html",
		pongo2.Context{"title": "Hello, World"})
}

func showSignInPage(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] show_sign_in")

	h.Render(c, "admin/signin.html",
		pongo2.Context{})
}

type SignInForm struct {
	Username string `form:"username" binding:"required,min=8,max=20"`
	Password string `form:"password" binding:"required,min=8,max=20"`
}

func signIn(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] sign_in")

	_, exists := c.Get(StashKeyAdmin)
	if exists {
		c.Redirect(http.StatusFound, "/admin/")
		return
	}

	var form SignInForm
	if c.Bind(&form) != nil {
		log.Infof(h.CTX(c), "[admin] form validation failed")
		// TODO showError
		return
	}

	admin, err := admins.FindByUsername(h.GOON(c), form.Username)
	if err != nil {
		log.Infof(h.CTX(c), "[admin] admin not found")
		// TODO showError not found
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(form.Password))
	if err != nil {
		log.Infof(h.CTX(c), "[admin] invalid password")
		// TODO showError unauthorized
		return
	}

	sess := sessions.Default(c)
	sess.Set(SessionKeyLoginAdminId, admin.AdminId)
	sess.Save()

	c.Redirect(http.StatusFound, "/admin/")
}

func signOut(c *gin.Context) {
	log.Infof(h.CTX(c), "[admin] sign_out")

	sess := sessions.Default(c)
	sess.Delete(SessionKeyLoginAdminId)
	sess.Save()

	c.Redirect(http.StatusFound, "/admin/signin")
}
