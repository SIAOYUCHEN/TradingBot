package domain

import (
	"context"

	dto "TradingBot/domain/dto"
)

type TradeRedisRepository interface {
	AddTrade(ctx context.Context, trade *dto.Trade, key string) error

	GetTrade(ctx context.Context, key string) ([]*dto.Trade, error)

	GetAllTrades(ctx context.Context) (map[string][]dto.Trade, error)

	ExistKey(ctx context.Context, key string) (int64, error)

	DeleteTrade(ctx context.Context, key string) (int64, error)

	MatchTrade(ctx context.Context, trade *dto.Trade, oppositeKey string, key string) ([]dto.MatchingTrade, error)
}
