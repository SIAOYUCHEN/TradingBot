package application

import (
	tradeInterface "TradingBot/domain/Interface"
	domain "TradingBot/domain/trade/deleteTrade"
	"context"
	"errors"
	"fmt"
)

type DeleteTradeHandler struct {
	TradeRepo tradeInterface.TradeRedisRepository
}

func NewDeleteTradeHandler(tradeRepo tradeInterface.TradeRedisRepository) *DeleteTradeHandler {
	return &DeleteTradeHandler{
		TradeRepo: tradeRepo,
	}
}

func (h *DeleteTradeHandler) Handle(ctx context.Context, command *domain.DeleteTradeCommand) (*domain.DeleteTradeResponse, error) {

	key := fmt.Sprintf("%s_%s", command.Market, command.Direction)

	exists, err := h.TradeRepo.ExistKey(ctx, key)
	if err != nil {
		return nil, errors.New("error checking if key exists")
	}

	if exists == 0 {
		return nil, errors.New("key not found")
	}

	_, err = h.TradeRepo.DeleteTrade(ctx, key)
	if err != nil {
		return nil, errors.New("error deleting trade")
	}

	return &domain.DeleteTradeResponse{
		Message: "Trades deleted successfully",
	}, nil
}
