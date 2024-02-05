package domain

import (
	dto "TradingBot/domain/dto"
	"context"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*dto.User, error)

	GetAllUsers(ctx context.Context) ([]dto.User, error)

	GetUser(ctx context.Context, id uint) (*dto.User, error)

	CreateUser(ctx context.Context, user *dto.User) (*dto.User, error)

	DeleteUser(ctx context.Context, user *dto.User) error

	UpdateUserEmail(ctx context.Context, id uint, email string) error
}
