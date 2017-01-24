package top_controller

import "github.com/gin-gonic/gin"

func Route(r *gin.RouterGroup) {
	r.GET("/", showIndex)
}
