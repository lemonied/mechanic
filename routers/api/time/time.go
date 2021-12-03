package apitime

import (
	"mechanic/models"
	"time"

	"github.com/gin-gonic/gin"
)

/*
Get get server time
*/
func Get(ctx *gin.Context) {
	ctx.JSON(models.MakeResponse(1, time.Now()))
}
