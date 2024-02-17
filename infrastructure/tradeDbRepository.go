package infrastructure

import (
	tradeInterface "TradingBot/domain/Interface"
	domain "TradingBot/domain/dto"
	"context"

	"gorm.io/gorm"
)

type GormTradeRepository struct {
	db *gorm.DB
}

func NewGormTradeRepository(db *gorm.DB) tradeInterface.TradeDbRepository {
	return &GormTradeRepository{db: db}
}

func (repo *GormTradeRepository) CreateOrder(ctx context.Context, trade *domain.Trades) error {
	if err := repo.db.WithContext(ctx).Create(trade).Error; err != nil {
		return err
	}

	return nil
}

func (repo *GormTradeRepository) CreateMatchTrade(ctx context.Context, matchTrade []*domain.MatchingTrades) error {
	if err := repo.db.WithContext(ctx).Create(matchTrade).Error; err != nil {
		return err
	}

	return nil
}
