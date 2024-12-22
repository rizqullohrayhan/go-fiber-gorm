package utils

import (
	"os"

	"github.com/rizqullorayhan/go-fiber-gorm/config"
)

func DeleteFile(file string) error {
	targetFile := config.ProjectRootPath + "/public/" + file
	return os.Remove(targetFile)
}