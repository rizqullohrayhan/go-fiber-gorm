package validators

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
)

func ValidateImage(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return errors.New("gagal membuka file")
	}
	defer src.Close()

	// Verifying if the file is a valid image by decoding it
	if _, _, err := image.Decode(src); err != nil {
		return errors.New("gambar tidak valid")
	}
	return nil
}