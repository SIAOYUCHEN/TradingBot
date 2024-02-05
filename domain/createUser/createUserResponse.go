package domain

type CreateResponse struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Email    *string `json:"email"`
}
