package domain

type UserResponse struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Email    *string `json:"email"`
}

type GetUserAllResponse struct {
	Users []UserResponse `json:"users"`
}
