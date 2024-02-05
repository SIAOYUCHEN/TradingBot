package domain

type GetUserQuery struct {
	Id uint `json:"id" validate:"required"`
}
