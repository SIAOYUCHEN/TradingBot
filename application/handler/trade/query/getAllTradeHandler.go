package application

import (
	tradeInterface "TradingBot/domain/Interface"
	domain "TradingBot/domain/trade/getAllTrade"
	"context"
)

type GetAllTradeHandler struct {
	TradeRepo tradeInterface.TradeRedisRepository
}

func NewGetAllTradeHandler(tradeRepo tradeInterface.TradeRedisRepository) *GetAllTradeHandler {
	return &GetAllTradeHandler{
		TradeRepo: tradeRepo,
	}
}

func (h *GetAllTradeHandler) Handle(ctx context.Context, query *domain.GetAllTradeQuery) (*domain.GetAllTradeResponse, error) {

	tradesWithKeys, err := h.TradeRepo.GetAllTrades(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.GetAllTradeResponse{
		TradesWithKeys: tradesWithKeys,
	}, nil
}
