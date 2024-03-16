package imageservice

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"

	"gocv.io/x/gocv"
)

type ImageService interface {
	ImageProcessesService(file multipart.File, x, y int) (bytes.Buffer, error)
}

type ImageServiceImpl struct {
}

func NewImageService() ImageService {
	return &ImageServiceImpl{}
}

func (i *ImageServiceImpl) ImageProcessesService(file multipart.File, x, y int) (bytes.Buffer, error) {
	fmt.Println("masuk service")
	var outputBuf bytes.Buffer

	// read file bytes
	imgData, err := io.ReadAll(file)
	if err != nil {
		return outputBuf, err
	}

	// decode to image
	img, err := gocv.IMDecode(imgData, gocv.IMReadColor)
	if err != nil {
		return outputBuf, err
	}

	resizedImg := gocv.NewMat()
	gocv.Resize(img, &resizedImg, image.Point{X: x, Y: y}, 0, 0, gocv.InterpolationDefault)

	// encode back to bytes
	buf := new(bytes.Buffer)
	im, _ := resizedImg.ToImage()
	err = jpeg.Encode(buf, im, &jpeg.Options{Quality: 75})
	if err != nil {
		return outputBuf, err
	}

	// encode to jpeg content type
	if err := jpeg.Encode(&outputBuf, im, nil); err != nil {
		return outputBuf, err
	}

	return outputBuf, err
}
