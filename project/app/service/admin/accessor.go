package admin

import (
	"app/model/admins"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func ADMIN(c *gin.Context) *admins.Admin {
	return c.MustGet(StashKeyAdmin).(*admins.Admin)
}
