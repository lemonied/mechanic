package routers

import (
	"mechanic/routers/api"

	"os"
	"path"

	"github.com/gin-gonic/gin"
)

/*
Register register router
*/
func Register(r *gin.Engine) {
	api.Register(r)
	var staticRoot, err = os.Getwd()
	if err == nil {
		staticRoot = path.Join(staticRoot, "./assets")
		r.Static("/static", path.Join(staticRoot, "./static"))
		r.StaticFile("/favicon.ico", path.Join(staticRoot, "./favicon.ico"))
		r.Use(func(c *gin.Context) {
			c.File(path.Join(staticRoot, "./index.html"))
		})
	}
}
