package entities

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Cover       string         `json:"cover"`
	IsAvailable bool           `json:"is_available" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	Categories        []Category          `json:"categories" gorm:"many2many:book_categories;"`
	BookBorrowHistory []BookBorrowHistory `json:"book_borrow_history" gorm:"foreignKey:BookID"`
}
