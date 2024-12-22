package dto

import (
	"time"
)

type CreateBookBorrowHistory struct {
	BookID     uint       `json:"book_id" form:"book_id" validate:"required"`
	UserID     uint       `json:"user_id" form:"user_id" validate:"required"`
	StatusID   uint       `json:"status_id" form:"status_id"`
	BorrowedAt time.Time  `json:"borrowed_at" form:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at" form:"returned_at"`
	Keterangan string     `json:"keterangan" form:"keterangan"`
}

type UpdateBookBorrowHistory struct {
	BookID     uint       `json:"book_id" form:"book_id" validate:"required"`
	UserID     uint       `json:"user_id" form:"user_id" validate:"required"`
	StatusID   uint       `json:"status_id" form:"status_id" validate:"required"`
	BorrowedAt time.Time  `json:"borrowed_at" form:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at" form:"returned_at"`
	Keterangan string     `json:"keterangan" form:"keterangan"`
}