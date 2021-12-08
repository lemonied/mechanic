package api

import (
	"mechanic/routers/api/graphics"
	"mechanic/routers/api/time"

	"github.com/gin-gonic/gin"
)

/*
Register register api
*/
func Register(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.GET("/time", apitime.Get)
	apiGroup.GET("/screenshot", graphics.Get)
	apiGroup.POST("/compare", graphics.ComparePost)
}
