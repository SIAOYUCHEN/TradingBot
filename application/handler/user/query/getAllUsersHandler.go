package application

import (
	domain "TradingBot/domain/getAllUsers"
	userInterface "TradingBot/domain/userInterface"
	"context"
)

type GetAllUsersHandler struct {
	UserRepo userInterface.UserRepository
}

func NewGetAllUsersHandler(userRepo userInterface.UserRepository) *GetAllUsersHandler {
	return &GetAllUsersHandler{
		UserRepo: userRepo,
	}
}

func (h *GetAllUsersHandler) Handle(ctx context.Context, queery *domain.GetAllUsersQuery) (*domain.GetUserAllResponse, error) {

	users, err := h.UserRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	usersResponse := make([]domain.UserResponse, len(users))
	for i, user := range users {
		usersResponse[i] = domain.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
	}

	return &domain.GetUserAllResponse{Users: usersResponse}, nil
}
