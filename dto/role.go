package dto

type GetRole struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateRole struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateRole struct {
	Name string `json:"name" form:"name" validate:"required"`
}
