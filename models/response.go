package models

import "github.com/gin-gonic/gin"

/*
MakeResponse makeResponse
*/
func MakeResponse(code int, data interface{}) (int, *gin.H) {
	return 200, &gin.H{
		"code": code,
		"data": data,
	}
}
