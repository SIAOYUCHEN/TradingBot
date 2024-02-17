package domain

import (
	dto "TradingBot/domain/dto"
	"context"
)

type TradeDbRepository interface {
	CreateOrder(ctx context.Context, trade *dto.Trades) error

	CreateMatchTrade(ctx context.Context, matchTrade []*dto.MatchingTrades) error
}
