package util

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

// Convert FileHeader to multipart.File
func ConvertFileHeaderToMultipartFile(filepath string) (*multipart.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	mtf := multipart.File(file)
	return &mtf, nil
}

// Delete file
func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

// Create file
func CreateFile(file *multipart.FileHeader) (string, error) {
	dir, _ := os.Getwd()

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	filepath := dir + file.Filename
	dst, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return filepath, nil
}

func ResizeAndSaveImage(filePath string, width, height int) error {
	// Open the existing image file
	src, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer src.Close()

	// Decode the image
	var img image.Image
	if strings.HasSuffix(strings.ToLower(filePath), ".png") {
		img, err = png.Decode(src)
		if err != nil {
			return err
		}
	} else {
		img, _, err = image.Decode(src)
		if err != nil {
			return err
		}
	}

	// Resize the image
	resizedImage := resize.Resize(uint(width), uint(height), img, resize.Lanczos2)

	// Open the file for writing (overwrite the original file)
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Save the resized image to the file
	err = jpeg.Encode(dst, resizedImage, nil)
	if err != nil {
		return err
	}

	return nil
}

// resize image file
func ResizeImg(img image.Image) image.Image {
	resizedImage := resize.Resize(250, 250, img, resize.Lanczos2)
	return resizedImage
}
