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
	if index != 15 {
		return "", errors.New("error")
	}
	imageType := base64String[11:index]

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64String[index+8:]))

	switch imageType {
	case "jpeg":
		jpegImg, err := jpeg.Decode(reader)
		if err != nil {
			return "", err
		}

		imageName := uuid.New().String() + ".jpeg"
		saveImage(jpegImg, imageName)
		return imageName, nil
	case "jpg":
		jpegImg, err := jpeg.Decode(reader)
		if err != nil {
			return "", err
		}

		imageName := uuid.New().String() + ".jpeg"
		saveImage(jpegImg, imageName)
		return imageName, nil

	case "png":
		pngImg, err := png.Decode(reader)
		if err != nil {
			return "", err
		}

		imageName := uuid.New().String() + ".png"
		saveImage(pngImg, imageName)
		return imageName, nil
	}

	return "", errors.New("error")
}

func saveImage(image image.Image, imageName string) {
	newImage := resize.Resize(200, 200, image, resize.Lanczos3)

	out, err := os.Create("./resources/avatar/" + imageName)
	PanicIfError(err)
	defer out.Close()

	jpeg.Encode(out, newImage, nil)
}
