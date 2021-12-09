package graphics

import (
	"image"
	"mechanic/pkg/picture"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

/*
ComparePost compare
*/
func ComparePost(ctx *gin.Context) (int, interface{}) {
	compareType, _ := ctx.GetQuery("type")
	file1, err := ctx.FormFile("file1")
	if err != nil {
		return 0, err.Error()
	}
	file2, err := ctx.FormFile("file2")
	if err != nil {
		return 0, err.Error()
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
	result, err := picture.Compare(images[0], images[1], compareType)
	if err != nil {
		return 0, err.Error()
	}
	return 1, result
}
