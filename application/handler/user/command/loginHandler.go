package application

import (
	userInterface "TradingBot/domain/Interface"
	dto "TradingBot/domain/dto"
	domain "TradingBot/domain/user/login"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginHandler struct {
	UserRepo userInterface.UserRepository
}

func NewLoginHandler(userRepo userInterface.UserRepository) *LoginHandler {
	return &LoginHandler{
		UserRepo: userRepo,
	}
}

func (h *LoginHandler) Handle(ctx context.Context, command *domain.LoginCommand) (*domain.LoginResponse, error) {
	user, err := h.UserRepo.GetUserByUsername(ctx, command.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("credentials are not correct")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(command.Password)); err != nil {
		return nil, errors.New("credentials are not correct")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &dto.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("TradingBot"))
	if err != nil {
		return nil, errors.New("could not create token")
	}

	return &domain.LoginResponse{Token: tokenString}, nil
}
