package application

import (
	tradeInterface "TradingBot/domain/Interface"
	dto "TradingBot/domain/dto"
	domain "TradingBot/domain/trade/createTrade"
	"context"
	"fmt"
)

type CreateTradeHandler struct {
	TradeRepo tradeInterface.TradeRedisRepository
}

func NewCreateTradeHandler(tradeRepo tradeInterface.TradeRedisRepository) *CreateTradeHandler {
	return &CreateTradeHandler{
		TradeRepo: tradeRepo,
	}
}

func (h *CreateTradeHandler) Handle(ctx context.Context, command *domain.CreateTradeCommand) (*domain.CreateTradeResponse, error) {

	key := fmt.Sprintf("%s_%s", command.Market, command.Direction)

	trade := &dto.Trade{
		Market:    command.Market,
		Price:     command.Price,
		Amount:    command.Amount,
		Direction: command.Direction,
	}

	if err := h.TradeRepo.AddTrade(ctx, trade, key); err != nil {
		return nil, err
	}

	return &domain.CreateTradeResponse{
		Message: "Trade created successfully",
	}, nil
}
