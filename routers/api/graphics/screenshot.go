package graphics

import (
	"mechanic/pkg/picture"

	"github.com/gin-gonic/gin"
)

/*
Get screenshot
*/
func Get(ctx *gin.Context) (int, interface{}) {
	ss := picture.Screenshot()
	b64, err := picture.ToBase64(ss.Image)
	if err == nil {
		return 1,  map[string]interface{}{
			"id": ss.ID,
			"base64": b64,
		}
	}
	return 0, err
}
