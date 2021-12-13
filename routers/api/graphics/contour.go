package graphics

import (
	"fmt"
	"image"
	"mechanic/pkg/picture"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

/*
ContourPost contour
*/
func ContourPost(ctx *gin.Context) (int, interface{}) {
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

	// 膨胀化
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(5, 5))
	defer kernel.Close()
	gocv.Erode(normalized, &normalized, kernel)

	// 查找轮廓
	contours := gocv.FindContours(normalized, gocv.RetrievalList, gocv.ChainApproxSimple)
	
	var dataURLs []string
	for i := 0; i < contours.Size(); i++ {
		contour := contours.At(i)
		area := gocv.ContourArea(contour)
		fmt.Printf("area: %v\n", area)
		rect := gocv.BoundingRect(contour)
		sliced := normalized.Region(rect)
		sliced = sliced.Clone()
		defer sliced.Close()
		resultImg, err := sliced.ToImage()
		if err != nil {
			panic(err)
		}
		dataURL, err := picture.ToBase64(resultImg)
		if err != nil {
			panic(err)
		}
		dataURLs = append(dataURLs, dataURL)

	}

	return 1, dataURLs
}
