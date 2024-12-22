package dto

import (
	// "time"

	"github.com/rizqullorayhan/go-fiber-gorm/entities"
)

type GetBook struct {
	ID                uint                         `json:"id"`
	Title             string                       `json:"title" form:"title"`
	Author            string                       `json:"author" form:"author"`
	Cover             string                       `json:"cover" form:"cover"`
	IsAvailable       bool                         `json:"is_available" form:"is_available"`
	Categories        []entities.Category          `json:"categories" form:"categories"`
	BookBorrowHistory []entities.BookBorrowHistory `json:"book_borrow_history"`
}

type CreateBook struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Author      string `json:"author" form:"author" validate:"required"`
	Cover       string `json:"cover" form:"cover"`
	IsAvailable bool   `json:"is_available" form:"is_available" validate:"required"`
	Categories  []int  `json:"categories" form:"categories" validate:"required,dive"`
}

type UpdateBook struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Author      string `json:"author" form:"author" validate:"required"`
	Cover       string `json:"cover" form:"cover"`
	IsAvailable bool   `json:"is_available" form:"is_available" validate:"required"`
	Categories  []int  `json:"categories" form:"categories" validate:"required"`
}
