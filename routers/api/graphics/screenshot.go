package graphics

import (
	"mechanic/models"
	"mechanic/pkg/robot"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

type screenshot struct {

}
/*
Get screenshot
*/
func Get(ctx *gin.Context) {
	ss := *robot.Screenshot()
	mat, err := gocv.NewMatFromBytes(ss.Width, ss.Height, gocv.MatTypeCV16S, ss.Data)
	if err == nil {
		ctx.JSON(models.MakeResponse(1, mat))
	} else {
		ctx.JSON(models.MakeResponse(0, err))
	}
}
