package application

import (
	userInterface "TradingBot/domain/Interface"
	domain "TradingBot/domain/user/getUser"
	"context"
	"errors"

	"gorm.io/gorm"
)

type GetUserHandler struct {
	UserRepo userInterface.UserRepository
}

func NewGetUserHandler(userRepo userInterface.UserRepository) *GetUserHandler {
	return &GetUserHandler{
		UserRepo: userRepo,
	}
}

func (h *GetUserHandler) Handle(ctx context.Context, query *domain.GetUserQuery) (*domain.GetUserResponse, error) {
	user, err := h.UserRepo.GetUser(ctx, query.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, errors.New("internal server error")
	}

	userResp := &domain.GetUserResponse{
		User: domain.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	return userResp, nil
}
