package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	RoleID    uint           `json:"role_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Role              Role                `json:"role"`
	BookBorrowHistory []BookBorrowHistory `json:"book_borrow_history" gorm:"foreignKey:UserID"`
}
