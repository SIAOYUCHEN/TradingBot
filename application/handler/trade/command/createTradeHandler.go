package application

import (
	tradeInterface "TradingBot/domain/Interface"
	dto "TradingBot/domain/dto"
	domain "TradingBot/domain/trade/createTrade"
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

type CreateTradeHandler struct {
	TradeRepo tradeInterface.TradeRedisRepository
	Node      *snowflake.Node
}

func NewCreateTradeHandler(tradeRepo tradeInterface.TradeRedisRepository, node *snowflake.Node) *CreateTradeHandler {
	return &CreateTradeHandler{
		TradeRepo: tradeRepo,
		Node:      node,
	}
}

func (h *CreateTradeHandler) Handle(ctx context.Context, command *domain.CreateTradeCommand) (*domain.CreateTradeResponse, error) {

	var oppositeDirection, currentDirection string

	if command.Direction == "Ask" {
		oppositeDirection = "Bid"
		currentDirection = "Ask"
	} else {
		oppositeDirection = "Ask"
		currentDirection = "Bid"
	}

	oppositeKey := fmt.Sprintf("%s_%s", command.Market, oppositeDirection)
	key := fmt.Sprintf("%s_%s", command.Market, currentDirection)

	id := h.Node.Generate().Int64()

	trade := &dto.Trade{
		Id:        id,
		Market:    command.Market,
		Price:     command.Price,
		Amount:    command.Amount,
		Direction: command.Direction,
		Timestamp: time.Now(),
	}

	if err := h.TradeRepo.MatchTrade(ctx, trade, oppositeKey, key); err != nil {
		return nil, err
	}

	return &domain.CreateTradeResponse{
		Message: "Trade created successfully",
	}, nil
}
