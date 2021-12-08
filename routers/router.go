package routers

import (
	"mechanic/models"
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
		r.NoRoute(func(c *gin.Context) {
			c.File(path.Join(staticRoot, "./index.html"))
		})
		r.Use(func(c *gin.Context) {
			c.Next()
			if length := len(c.Errors); length > 0 {
				e := c.Errors[length-1]
				err := e.Err
				if err != nil {
					c.JSON(models.MakeResponse(0, err))
				}
			}
		})
	}
}
