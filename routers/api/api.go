package api

import (
	"mechanic/pkg/utils"
	"mechanic/routers/api/graphics"
	"mechanic/routers/api/time"

	"github.com/gin-gonic/gin"
)

/*
Register register api
*/
func Register(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.GET("/time", utils.MakeJSON(apitime.Get))
	apiGroup.GET("/screenshot", utils.MakeJSON(graphics.Get))
	apiGroup.POST("/compare", utils.MakeJSON(graphics.ComparePost))
	apiGroup.POST("/find", utils.MakeJSON(graphics.SearchPost))
}
