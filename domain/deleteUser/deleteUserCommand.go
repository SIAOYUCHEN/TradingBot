package domain

type DeleteUserCommand struct {
	Id uint `json:"id" validate:"required"`
}
