package infrastructure

import (
	domain "TradingBot/domain/dto"
	userInterface "TradingBot/domain/userInterface"
	"context"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) userInterface.UserRepository {
	return &GormUserRepository{db: db}
}

func (repo *GormUserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User

	result := repo.db.WithContext(ctx).Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *GormUserRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	result := repo.db.WithContext(ctx).Select("ID", "Username", "Email").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repo *GormUserRepository) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User

	result := repo.db.WithContext(ctx).Select("ID", "Username", "Email").First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *GormUserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *GormUserRepository) DeleteUser(ctx context.Context, user *domain.User) error {
	if result := repo.db.Delete(user).Error; result != nil {
		return result
	}

	return nil
}

func (repo *GormUserRepository) UpdateUserEmail(ctx context.Context, id uint, email string) error {
	result := repo.db.Model(&domain.User{}).Where("id = ?", id).Update("email", email)

	return result.Error
}
