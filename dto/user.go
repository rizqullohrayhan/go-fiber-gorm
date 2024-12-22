package dto

import "github.com/rizqullorayhan/go-fiber-gorm/entities"

// type CreateUser struct {
// 	Name     string `json:"name" validate:"required"`
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required"`
// 	Address  string `json:"address" validate:"required"`
// 	Phone    string `json:"phone" validate:"required"`
// }

type GetUser struct {
	ID                uint                         `json:"id"`
	Name              string                       `json:"name"`
	Email             string                       `json:"email"`
	Address           string                       `json:"address"`
	Phone             string                       `json:"phone"`
	Role              string                       `json:"role"`
	BookBorrowHistory []entities.BookBorrowHistory `json:"book_borrow_history" gorm:"foreignKey:UserID"`
}

type UpdateUser struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

type UpdateUserRole struct {
	RoleID uint `json:"role_id" validate:"required"`
}

type UpdateUserEmail struct {
	Email string `json:"email" validate:"required,email"`
}
