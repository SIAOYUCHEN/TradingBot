package domain

type UpdateUserEmailCommand struct {
	Id    uint   `json:"id" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type Email struct {
	Email string `json:"email" example:"user@example.com"`
}
