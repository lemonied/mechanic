package graphics

import (
	"image"
	"mechanic/pkg/picture"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

/*
SearchPost find image
*/
func SearchPost(ctx *gin.Context) (int, interface{}) {
	var matchMode gocv.TemplateMatchMode
	mode, _ := ctx.GetQuery("mode")

	switch mode {
		case "TM_SQDIFF":
			matchMode = gocv.TmSqdiff
		case "TM_SQDIFF_NORMED":
			matchMode = gocv.TmSqdiffNormed
		case "TM_CCORR":
			matchMode = gocv.TmCcorr
		case "TM_CCORR_NORMED":
			matchMode = gocv.TmCcorrNormed
		case "TM_CCOEFF":
			matchMode = gocv.TmCcoeff
		case "TM_CCOEFF_NORMED":
			matchMode = gocv.TmCcoeffNormed
		default:
			matchMode = gocv.TmSqdiff
	}
	
	file1, err := ctx.FormFile("file1")
	if err != nil {
		return 0, err
	}
	file2, err := ctx.FormFile("file2")
	if err != nil {
		return 0, err
	}
	files := [2]*multipart.FileHeader{file1, file2}
	
	var images []image.Image
	
	for _, file := range files {
		reader, err := file.Open()
		defer reader.Close()
		if err != nil {
			return 0, err.Error()
		}
		image, _, err := image.Decode(reader)
		if err != nil {
			return 0, err.Error()
		}
		images = append(images, image)
	}
	result, err := picture.FindImage(images[0], images[1], matchMode)
	if err != nil {
		return 0, err.Error()
	}
	return 1, result
}
