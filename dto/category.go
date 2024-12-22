package dto

import (
	"time"

	"github.com/rizqullorayhan/go-fiber-gorm/entities"
)

type GetCategory struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	Books     []entities.Book `json:"books"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type CreateCategory struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateCategory struct {
	Name string `json:"name" form:"name" validate:"required"`
}
