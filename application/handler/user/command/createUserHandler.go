package application

import (
	domain "TradingBot/domain/createUser"
	dto "TradingBot/domain/dto"
	userInterface "TradingBot/domain/userInterface"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserHandler struct {
	UserRepo userInterface.UserRepository
}

func NewCreateUserHandler(userRepo userInterface.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		UserRepo: userRepo,
	}
}

func (h *CreateUserHandler) Handle(ctx context.Context, command *domain.CreateUserCommand) (*domain.CreateResponse, error) {

	existingUser, _ := h.UserRepo.GetUserByUsername(ctx, command.Username)

	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(command.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("password encryption failed")
	}

	newUser := dto.User{
		Username: command.Username,
		Password: string(hashedPassword),
		Email:    command.Email,
	}

	createdUser, err := h.UserRepo.CreateUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	return &domain.CreateResponse{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}, nil
}
