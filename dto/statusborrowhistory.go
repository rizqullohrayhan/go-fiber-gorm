package dto

type CreateStatusBorrowHistory struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateStatusBorrowHistory struct {
	Name string `json:"name" form:"name" validate:"required"`
}