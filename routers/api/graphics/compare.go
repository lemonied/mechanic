package graphics

import (
	"image"
	"mechanic/models"
	"mechanic/pkg/picture"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

/*
ComparePost compare
*/
func ComparePost(ctx *gin.Context) {
	compareType, compareTypeBool := ctx.GetQuery("type")
	if !compareTypeBool {
		compareType = "phash"
	}
	file1, err := ctx.FormFile("file1")
	if err != nil {
		ctx.JSON(models.MakeResponse(0, err))
		return
	}
	file2, err := ctx.FormFile("file2")
	if err != nil {
		ctx.JSON(models.MakeResponse(0, err))
		return
	}
	files := [2]*multipart.FileHeader{file1, file2}
	
	var images []image.Image
	
	for _, file := range files {
		reader, err := file.Open()
		if err != nil {
			ctx.JSON(models.MakeResponse(0, err))
			return
		}
		image, _, err := image.Decode(reader)
		if err != nil {
			ctx.JSON(models.MakeResponse(0, err))
			return
		}
		images = append(images, image)
	}
	result, err := picture.Compare(images[0], images[1], compareType)
	if err != nil {
		ctx.JSON(models.MakeResponse(0, err))
		return
	}
	ctx.JSON(models.MakeResponse(1, result))
}
