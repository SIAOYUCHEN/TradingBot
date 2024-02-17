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
	TradeDb   tradeInterface.TradeDbRepository
}

func NewCreateTradeHandler(tradeRepo tradeInterface.TradeRedisRepository, node *snowflake.Node, tradeDb tradeInterface.TradeDbRepository) *CreateTradeHandler {
	return &CreateTradeHandler{
		TradeRepo: tradeRepo,
		Node:      node,
		TradeDb:   tradeDb,
	}
}

func (h *CreateTradeHandler) Handle(ctx context.Context, command *domain.CreateTradeCommand) (*domain.CreateTradeResponse, error) {

	var oppositeDirection, currentDirection string

	var matchingTradesSlice []*dto.MatchingTrades

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

	trades := &dto.Trades{
		Id:        id,
		Market:    command.Market,
		Price:     command.Price,
		Amount:    command.Amount,
		Direction: command.Direction,
		Timestamp: time.Now(),
	}

	if err := h.TradeDb.CreateOrder(ctx, trades); err != nil {
		return nil, err
	}

	if err := h.TradeRepo.AddTrade(ctx, trade, "Order"); err != nil {
		return nil, err
	}

	matchedTrades, err := h.TradeRepo.MatchTrade(ctx, trade, oppositeKey, key)
	if err != nil {
		return nil, err
	}

	for _, mt := range matchedTrades {
		matchingTrade := &dto.MatchingTrades{
			Id:        h.Node.Generate().Int64(),
			AskId:     mt.AskId,
			BidId:     mt.BidId,
			Market:    mt.Market,
			Price:     mt.Price,
			Amount:    mt.Amount,
			Timestamp: mt.Timestamp,
		}
		matchingTradesSlice = append(matchingTradesSlice, matchingTrade)
	}

	if len(matchingTradesSlice) > 0 {
		if err := h.TradeDb.CreateMatchTrade(ctx, matchingTradesSlice); err != nil {
			return nil, err
		}
	}

	return &domain.CreateTradeResponse{
		Message: "Trade created successfully",
	}, nil
}
