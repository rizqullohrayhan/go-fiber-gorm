package entities

import (
	"time"

	"gorm.io/gorm"
)

type BookBorrowHistory struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	BookID     uint       `json:"book_id"`
	UserID     uint       `json:"user_id"`
	StatusID   uint       `json:"status_id"`
	BorrowedAt time.Time  `json:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at"`
	Keterangan string     `json:"keterangan"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Book   Book                `json:"book" gorm:"foreignKey:BookID"`
	User   User                `json:"user" gorm:"foreignKey:UserID"`
	Status StatusBorrowHistory `json:"status" gorm:"foreignKey:StatusID"`
}
