package graphics

import (
	"image"
	"mechanic/pkg/picture"
	"strconv"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

/*
BilateralFilter bilateral filter
*/
func BilateralFilter(ctx *gin.Context) (int, interface{}) {
	file, err := ctx.FormFile("file")
	diameter, err := strconv.Atoi(ctx.Query("diameter"))
	if err != nil {
		panic(err)
	}
	sigmaColor, err := strconv.ParseFloat(ctx.Query("sigmaColor"), 32)
	if err != nil {
		panic(err)
	}
	sigmaSpace, err := strconv.ParseFloat(ctx.Query("sigmaSpace"), 32)
	if err != nil {
		panic(err)
	}
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
	src, err := gocv.ImageToMatRGB(img)
	if err != nil {
		panic(err)
	}
	defer src.Close()
	target := gocv.NewMat()
	defer target.Close()
	gocv.BilateralFilter(src, &target, diameter, sigmaColor, sigmaSpace)
	targetImg, err := target.ToImage()
	if err != nil {
		panic(err)
	}
	dataURL, err := picture.ToBase64(targetImg)
	if err != nil {
		panic(err)
	}
	return 1, dataURL
}
