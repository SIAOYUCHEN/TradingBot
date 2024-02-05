package application

import (
	domain "TradingBot/domain/deleteUser"
	userInterface "TradingBot/domain/userInterface"
	"context"
	"errors"

	"gorm.io/gorm"
)

type DeleteUserHandler struct {
	UserRepo userInterface.UserRepository
}

func NewDeleteUserHandler(userRepo userInterface.UserRepository) *DeleteUserHandler {
	return &DeleteUserHandler{
		UserRepo: userRepo,
	}
}

func (h *DeleteUserHandler) Handle(ctx context.Context, command *domain.DeleteUserCommand) (*domain.DeleteUserResponse, error) {

	user, err := h.UserRepo.GetUser(ctx, command.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}

		return nil, errors.New("Internal server error")
	}

	if err = h.UserRepo.DeleteUser(ctx, user); err != nil {
		return nil, errors.New("Could not delete user")
	}

	return &domain.DeleteUserResponse{Message: "User deleted successfully"}, nil
}
