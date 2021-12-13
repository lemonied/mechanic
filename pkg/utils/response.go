package utils

import (

	"github.com/gin-gonic/gin"
)

/*
MakeJSONCallback json callback
*/
type MakeJSONCallback = func(ctx *gin.Context) (int, interface{})

/*
MakeJSON make json response
*/
func MakeJSON(fn MakeJSONCallback) (func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if _, ok := err.(error); ok {
				ctx.JSON(200, &gin.H{
					"code": 0,
					"error": err.(error).Error(),
				})
			}
		}()
		code, data := fn(ctx)
		ctx.JSON(200, &gin.H{
			"code": code,
			"data": data,
		})
	}
}
