package application

import (
	userInterface "TradingBot/domain/Interface"
	domain "TradingBot/domain/user/updateUserEmail"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UpdateUserEmailHandler struct {
	UserRepo userInterface.UserRepository
}

func NewUpdateUserEmailHandler(userRepo userInterface.UserRepository) *UpdateUserEmailHandler {
	return &UpdateUserEmailHandler{
		UserRepo: userRepo,
	}
}

func (h *UpdateUserEmailHandler) Handle(ctx context.Context, command *domain.UpdateUserEmailCommand) (*domain.UpdateUserEmailResponse, error) {

	user, err := h.UserRepo.GetUser(ctx, command.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}

		return nil, errors.New("Internal server error")
	}

	err = h.UserRepo.UpdateUserEmail(ctx, user.ID, command.Email)
	if err != nil {
		return nil, errors.New("Could not update user email")
	}

	return &domain.UpdateUserEmailResponse{Message: "User email updated"}, nil
}
