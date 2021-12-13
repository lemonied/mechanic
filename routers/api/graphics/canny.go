package graphics

import (
	"image"
	"mechanic/pkg/picture"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

/*
Canny canny
*/
func Canny(ctx *gin.Context) (int, interface{}) {
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

	sourceMat, err := gocv.ImageToMatRGB(img)
	defer sourceMat.Close()
	if err != nil {
		panic(err)
	}
	
	// 双边滤波
	filtered := gocv.NewMat()
	defer filtered.Close()
	gocv.BilateralFilter(sourceMat, &filtered, 80, 150, 150)

	// 灰度化
	grayscale := gocv.NewMat()
	defer grayscale.Close()
	gocv.CvtColor(filtered, &grayscale, gocv.ColorRGBToGray)

	// 二值化
	normalized := gocv.NewMat()
	defer normalized.Close()
	gocv.Threshold(grayscale, &normalized, 150, 255, gocv.ThresholdBinaryInv)

	// 边缘检测
	edge := gocv.NewMat()
	defer edge.Close()
	gocv.Canny(normalized, &edge, 100, 200)

	targetImg, err := edge.ToImage()
	if err != nil {
		panic(err)
	}
	dataURL, err := picture.ToBase64(targetImg)
	if err != nil {
		panic(err)
	}
	return 1, dataURL
}
