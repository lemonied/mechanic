package apitime

import (
	"time"

	"github.com/gin-gonic/gin"
)

/*
Get get server time
*/
func Get(ctx *gin.Context) (int, interface{}) {
	return 1, time.Now()
}
