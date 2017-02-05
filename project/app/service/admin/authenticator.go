package admin

import (
	"net/http"

	ch "app/controller_helper"
	"app/model/admins"

	"github.com/gin-contrib/sessions"
	"google.golang.org/appengine/log"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func authenticator(fallbackPath string,
	exceptionPaths []string) gin.HandlerFunc {

	return func(c *gin.Context) {

		if !needAuth(c.Request.URL.Path, exceptionPaths) {
			c.Next()
			return
		}

		sess := sessions.Default(c)

		value := sess.Get(SessionKeyLoginAdminId)
		admin_id, ok := value.(int64)

		if !ok {
			// TODO save destination

			log.Infof(ch.CTX(c), "[admin] session not found")

			c.Redirect(http.StatusFound, fallbackPath)
			c.Abort()
			return
		}

		admin, err := admins.FindById(ch.GOON(c), admin_id)
		if err != nil {
			log.Infof(ch.CTX(c), "[admin] invalid session")

			sess.Delete(SessionKeyLoginAdminId)
			sess.Save()

			c.Redirect(http.StatusFound, fallbackPath)
			c.Abort()
			return
		}

		c.Set(StashKeyAdmin, admin)
		c.Next()
	}
}

func needAuth(path string, exceptions []string) bool {
	for _, exp := range exceptions {
		if path == exp {
			return false
		}
	}
	return true
}
