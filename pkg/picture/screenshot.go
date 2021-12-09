package picture

import (
	"image"
	"mechanic/pkg/utils"

	"github.com/go-vgo/robotgo"
)

/*
Shot screenshot
*/
type Shot struct {
	ID string
	Image image.Image
}
/*
Screenshot get screenshot
args[0] -> x  args[1] -> y  args[2] -> width args[3] -> height
*/
func Screenshot(args ...int) Shot {
  ss := robotgo.CaptureImg(args...)
  return Shot{
		ID: utils.RandomStr(8),
		Image: ss,
	}
}
