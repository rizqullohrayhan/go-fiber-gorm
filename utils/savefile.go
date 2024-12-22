package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/rizqullorayhan/go-fiber-gorm/config"
)

func SaveFile(file *multipart.FileHeader, folder string, fileName string) error {
	// Create folder if it does not exist
	destFolder := config.ProjectRootPath+"/public/"+folder
	if err := os.MkdirAll(destFolder, os.ModePerm); err != nil {
		return errors.New("gagal membuat folder penyimpanan")
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return errors.New("gagal membuka file yang diunggah")
	}
	defer src.Close()

	// Create the destination file
	destPath := filepath.Join(destFolder, fileName)
	dest, err := os.Create(destPath)
	if err != nil {
		return errors.New("gagal membuat file tujuan")
	}
	defer dest.Close()

	// Copy the content of the uploaded file to the destination file
	if _, err := io.Copy(dest, src); err != nil {
		return errors.New("gagal menyimpan file")
	}

	return nil
}