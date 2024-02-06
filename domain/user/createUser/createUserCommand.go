package domain

type CreateUserCommand struct {
	Username string  `json:"username" validate:"required"`
	Password string  `json:"password" validate:"required"`
	Email    *string `json:"email" validate:"omitempty,email"`
}
