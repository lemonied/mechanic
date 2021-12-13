package graphics

import (
	"image"
	"mechanic/pkg/picture"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

/*
NormalizedPost normalized
*/
func NormalizedPost(ctx *gin.Context) (int, interface{}) {
	file, err := ctx.FormFile("file")
	if err != nil {
		panic(err)
	}
	reader, err := file.Open()
	defer reader.Close()
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	// Image 转 Mat
	sourceMat, err := gocv.ImageToMatRGB(img)
	defer sourceMat.Close()
	if err != nil {
		panic(err)
	}

	// 灰度化
	grayscale := gocv.NewMat()
	defer grayscale.Close()
	gocv.CvtColor(sourceMat, &grayscale, gocv.ColorRGBToGray)
	
	// 二值化
	normalized := gocv.NewMat()
	defer normalized.Close()
	gocv.Threshold(grayscale, &normalized, 150, 255, gocv.ThresholdBinaryInv)

	resultImg, err := normalized.ToImage()
	if err != nil {
		panic(err)
	}
	dataURL, err := picture.ToBase64(resultImg)
	if err != nil {
		panic(err)
	}
	return 1, dataURL
}
