package graphics

import (
	"mechanic/models"
	"mechanic/pkg/picture"
	"mechanic/pkg/robot"

	"github.com/gin-gonic/gin"
)

/*
Get screenshot
*/
func Get(ctx *gin.Context) {
	ss := robot.Screenshot()
	b64, err := picture.ToBase64(ss.Image)
	if err == nil {
		ctx.JSON(models.MakeResponse(1, map[string]interface{}{
			"id": ss.ID,
			"base64": b64,
		}))
	} else {
		ctx.JSON(models.MakeResponse(0, err))
	}
}
