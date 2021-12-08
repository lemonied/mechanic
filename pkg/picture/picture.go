package picture

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

/*
ToBase64 bitmap to image base64
*/
func ToBase64(image image.Image) (string, error) {
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, image)
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), err
}

/*
Compare compare two images similarity 0 ~ 1
hashType: phash | average | blockmean0 | blockmean1 | colormoment | marrhildreth | radialvariance
*/
func Compare(source image.Image, target image.Image, hashType string) (float64, error) {
	var hash contrib.ImgHashBase

	switch hashType {
		case "phash":
			hash = contrib.PHash{}
		case "average":
			hash = contrib.AverageHash{}
		case "blockmean0":
			hash = contrib.BlockMeanHash{}
		case "blockmean1":
			hash = contrib.BlockMeanHash{ Mode: contrib.BlockMeanHashMode1 }
		case "colormoment":
			hash = contrib.ColorMomentHash{}
		case "marrhildreth":
			hash = contrib.NewMarrHildrethHash()
		case "radialvariance":
			hash = contrib.NewRadialVarianceHash()
		default:
			hash = contrib.PHash{}
	}

	mat1, err1 := gocv.ImageToMatRGBA(source)
	defer mat1.Close()
	if err1 != nil {
		return 0, err1
	}
	
	mat2, err2 := gocv.ImageToMatRGBA(target)
	defer mat2.Close()
	if err2 != nil {
		return 0, err2
	}

	result1 := gocv.NewMat()
	defer result1.Close()
	result2 := gocv.NewMat()
	defer result2.Close()
	hash.Compute(mat1, &result1)
	hash.Compute(mat2, &result2)

	similar := hash.Compare(result1, result2)

	return similar, nil

}
