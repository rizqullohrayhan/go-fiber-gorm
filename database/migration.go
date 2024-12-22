package database

import (
	"fmt"

	"github.com/rizqullorayhan/go-fiber-gorm/entities"
)

func MigrateInit() error {
	err := DB.AutoMigrate(
		&entities.Role{},
		&entities.User{},
		&entities.Category{},
		&entities.Book{},
		&entities.StatusBorrowHistory{},
		&entities.BookBorrowHistory{},
	)
	if err != nil {
		return fmt.Errorf("terjadi kesalahan dalam migration: %v", err)
	}
	return nil
}