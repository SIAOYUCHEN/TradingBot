package application

import (
	tradeInterface "TradingBot/domain/Interface"
	dto "TradingBot/domain/dto"
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
)

type KafkaTradeMessageHandler struct {
	TradeDb       tradeInterface.TradeDbRepository
	TradeRepo     tradeInterface.TradeRedisRepository
	KafkaProducer tradeInterface.KafkaProducerRepository
}

func NewKafkaTradeMessageHandler(tradeDb tradeInterface.TradeDbRepository, tradeRepo tradeInterface.TradeRedisRepository,
	kafkaProducer tradeInterface.KafkaProducerRepository) *KafkaTradeMessageHandler {
	return &KafkaTradeMessageHandler{
		TradeDb:       tradeDb,
		TradeRepo:     tradeRepo,
		KafkaProducer: kafkaProducer,
	}
}

func (h *KafkaTradeMessageHandler) Setup(sarama.ConsumerGroupSession) error {

	return nil
}

func (h *KafkaTradeMessageHandler) Cleanup(sarama.ConsumerGroupSession) error {

	return nil
}

func (h *KafkaTradeMessageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {

		switch message.Topic {
		case "trades":
			h.handleTradesMessage(message)
		case "other-topic":
			//h.handleOtherTopicMessage(message)
		}

		session.MarkMessage(message, "")
	}

	return nil
}

func (h *KafkaTradeMessageHandler) handleTradesMessage(message *sarama.ConsumerMessage) {
	var tradeMessage dto.Trade
	if err := json.Unmarshal(message.Value, &tradeMessage); err != nil {

		return
	}

	trades := dto.Trades{
		Id:        tradeMessage.Id,
		Market:    tradeMessage.Market,
		Price:     tradeMessage.Price,
		Amount:    tradeMessage.Amount,
		Direction: tradeMessage.Direction,
		Timestamp: tradeMessage.Timestamp,
	}

	trade := dto.Trade{
		Id:        tradeMessage.Id,
		Market:    tradeMessage.Market,
		Price:     tradeMessage.Price,
		Amount:    tradeMessage.Amount,
		Direction: tradeMessage.Direction,
		Timestamp: tradeMessage.Timestamp,
	}

	if err := h.TradeDb.CreateOrder(context.Background(), &trades); err != nil {

		return
	}

	if err := h.TradeRepo.AddTrade(context.Background(), &trade, "Order"); err != nil {

		return
	}

	h.KafkaProducer.SendTradeToKafkaSpecifyTopic(&trade, "transactions")
}
