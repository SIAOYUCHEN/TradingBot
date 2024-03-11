package application

import (
	tradeInterface "TradingBot/domain/Interface"
	dto "TradingBot/domain/dto"
	"context"
	"encoding/json"

	"fmt"

	"log"

	"github.com/IBM/sarama"
	"github.com/bwmarrin/snowflake"
)

type KafkaTradeMessageHandler struct {
	TradeDb       tradeInterface.TradeDbRepository
	TradeRepo     tradeInterface.TradeRedisRepository
	KafkaProducer tradeInterface.KafkaProducerRepository
	Node          *snowflake.Node
}

func NewKafkaTradeMessageHandler(tradeDb tradeInterface.TradeDbRepository, tradeRepo tradeInterface.TradeRedisRepository,
	kafkaProducer tradeInterface.KafkaProducerRepository, node *snowflake.Node) *KafkaTradeMessageHandler {
	return &KafkaTradeMessageHandler{
		TradeDb:       tradeDb,
		TradeRepo:     tradeRepo,
		KafkaProducer: kafkaProducer,
		Node:          node,
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
		log.Printf("Message claimed: topic = %s, partition = %d, offset = %d\n", message.Topic, message.Partition, message.Offset)

		switch message.Topic {
		case "trades":
			log.Println("Handling trades message")
			h.handleTradesMessage(message)
		case "matching":
			log.Println("Handling matching message")
			h.handleMatchingMessage(message)
		case "transactions":
			log.Println("Handling transaction message")
			h.handleTransactionMessage(message)
		default:
			log.Printf("Unknown topic: %s\n", message.Topic)
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

	h.KafkaProducer.SendTradeToKafkaSpecifyTopic(&trade, "matching")
}

func (h *KafkaTradeMessageHandler) handleMatchingMessage(message *sarama.ConsumerMessage) {

	var tradeMessage dto.Trade

	if err := json.Unmarshal(message.Value, &tradeMessage); err != nil {

		return
	}

	var oppositeDirection, currentDirection string

	if tradeMessage.Direction == "Ask" {
		oppositeDirection = "Bid"
		currentDirection = "Ask"
	} else {
		oppositeDirection = "Ask"
		currentDirection = "Bid"
	}

	oppositeKey := fmt.Sprintf("%s_%s", tradeMessage.Market, oppositeDirection)
	key := fmt.Sprintf("%s_%s", tradeMessage.Market, currentDirection)

	matchedTrades, err := h.TradeRepo.MatchTrade(context.Background(), &tradeMessage, oppositeKey, key)

	if err != nil {
		return
	}

	if len(matchedTrades) > 0 {
		if err := h.KafkaProducer.SendMatchTradeToKafkaSpecifyTopic(matchedTrades, "transactions"); err != nil {
			log.Printf("Error sending matched trades to Kafka: %v", err)
		}
	}
}

func (h *KafkaTradeMessageHandler) handleTransactionMessage(message *sarama.ConsumerMessage) {

	var matchingTradesSlice []*dto.MatchingTrades

	if err := json.Unmarshal(message.Value, &matchingTradesSlice); err != nil {

		return
	}

	if len(matchingTradesSlice) == 0 {

		return
	}

	for i, mt := range matchingTradesSlice {
		mt.Id = h.Node.Generate().Int64()
		matchingTradesSlice[i] = mt
	}

	if err := h.TradeDb.CreateMatchTrade(context.Background(), matchingTradesSlice); err != nil {

		return
	}
}
