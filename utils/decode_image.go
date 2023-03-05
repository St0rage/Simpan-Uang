package utils

import (
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func DecodeImage(base64String string) (string, error) {
	index := strings.Index(base64String, ";base64,")
	imageType := base64String[11:index]

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64String[index+8:]))

	switch imageType {
	case "jpeg":
		jpegImg, err := jpeg.Decode(reader)
		PanicIfError(err)

		imageName := uuid.New().String() + ".jpeg"
		saveImage(jpegImg, imageName)
		return imageName, nil
	case "jpg":
		jpegImg, err := jpeg.Decode(reader)
		PanicIfError(err)

		imageName := uuid.New().String() + ".jpeg"
		saveImage(jpegImg, imageName)
		return imageName, nil

	case "png":
		pngImg, err := png.Decode(reader)
		PanicIfError(err)

		imageName := uuid.New().String() + ".png"
		saveImage(pngImg, imageName)
		return imageName, nil
	}

	return "", errors.New("error")
}

func saveImage(image image.Image, imageName string) {
	newImage := resize.Resize(200, 200, image, resize.Lanczos3)

	out, err := os.Create("./resources/avatar/" + imageName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	jpeg.Encode(out, newImage, nil)
}
