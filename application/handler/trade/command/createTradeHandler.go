package application

import (
	tradeInterface "TradingBot/domain/Interface"
	dto "TradingBot/domain/dto"
	domain "TradingBot/domain/trade/createTrade"
	"context"
	"time"

	"github.com/bwmarrin/snowflake"
)

type CreateTradeHandler struct {
	TradeRepo     tradeInterface.TradeRedisRepository
	Node          *snowflake.Node
	TradeDb       tradeInterface.TradeDbRepository
	KafkaProducer tradeInterface.KafkaProducerRepository
}

func NewCreateTradeHandler(tradeRepo tradeInterface.TradeRedisRepository, node *snowflake.Node,
	tradeDb tradeInterface.TradeDbRepository, kafkaProducer tradeInterface.KafkaProducerRepository) *CreateTradeHandler {
	return &CreateTradeHandler{
		TradeRepo:     tradeRepo,
		Node:          node,
		TradeDb:       tradeDb,
		KafkaProducer: kafkaProducer,
	}
}

func (h *CreateTradeHandler) Handle(ctx context.Context, command *domain.CreateTradeCommand) (*domain.CreateTradeResponse, error) {

	// var oppositeDirection, currentDirection string

	// var matchingTradesSlice []*dto.MatchingTrades

	// if command.Direction == "Ask" {
	// 	oppositeDirection = "Bid"
	// 	currentDirection = "Ask"
	// } else {
	// 	oppositeDirection = "Ask"
	// 	currentDirection = "Bid"
	// }

	// oppositeKey := fmt.Sprintf("%s_%s", command.Market, oppositeDirection)
	// key := fmt.Sprintf("%s_%s", command.Market, currentDirection)

	id := h.Node.Generate().Int64()

	trade := &dto.Trade{
		Id:        id,
		Market:    command.Market,
		Price:     command.Price,
		Amount:    command.Amount,
		Direction: command.Direction,
		Timestamp: time.Now(),
	}

	// trades := &dto.Trades{
	// 	Id:        id,
	// 	Market:    command.Market,
	// 	Price:     command.Price,
	// 	Amount:    command.Amount,
	// 	Direction: command.Direction,
	// 	Timestamp: time.Now(),
	// }

	if err := h.KafkaProducer.SendTradeToKafkaSpecifyTopic(trade, "trades"); err != nil {
		return nil, err
	}

	return &domain.CreateTradeResponse{
		Message: "Trade created successfully",
	}, nil
}
