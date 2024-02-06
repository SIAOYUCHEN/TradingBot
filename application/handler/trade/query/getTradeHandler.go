package application

import (
	tradeInterface "TradingBot/domain/Interface"
	dto "TradingBot/domain/dto"
	domain "TradingBot/domain/trade/getTrade"
	"context"
	"fmt"
)

type GetTradeHandler struct {
	TradeRepo tradeInterface.TradeRedisRepository
}

func NewGetTradeHandler(tradeRepo tradeInterface.TradeRedisRepository) *GetTradeHandler {
	return &GetTradeHandler{
		TradeRepo: tradeRepo,
	}
}

func (h *GetTradeHandler) Handle(ctx context.Context, query *domain.GetTradeQuery) (*domain.GetTradeResponse, error) {

	key := fmt.Sprintf("%s_%s", query.Market, query.Direction)

	trades, err := h.TradeRepo.GetTrade(ctx, key)
	if err != nil {
		return nil, err
	}

	var tradesDTO []dto.Trade
	for _, trade := range trades {
		tradesDTO = append(tradesDTO, dto.Trade{
			Market:    trade.Market,
			Price:     trade.Price,
			Amount:    trade.Amount,
			Direction: trade.Direction,
		})
	}

	return &domain.GetTradeResponse{
		Trades: tradesDTO,
	}, nil
}
