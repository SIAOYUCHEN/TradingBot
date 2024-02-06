package domain

type UserResponse struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Email    *string `json:"email"`
}

type GetUserResponse struct {
	User UserResponse `json:"users"`
}
