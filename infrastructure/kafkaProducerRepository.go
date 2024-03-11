package infrastructure

import (
	"github.com/IBM/sarama"

	kafkaInterface "TradingBot/domain/Interface"

	domain "TradingBot/domain/dto"
	"encoding/json"
	"fmt"
)

type KafkaProducerRepository struct {
	Producer sarama.SyncProducer
}

func NewKafkaProducerRepository(brokers []string) kafkaInterface.KafkaProducerRepository {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	return &KafkaProducerRepository{
		Producer: producer,
	}
}

func (k *KafkaProducerRepository) SendMessage(topic string, key, message []byte) error {
	producerMessage := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := k.Producer.SendMessage(producerMessage)
	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaProducerRepository) SendTradeToKafkaSpecifyTopic(trade *domain.Trade, topic string) error {
	tradeJSON, err := json.Marshal(trade)
	if err != nil {
		return err
	}

	tradeID := []byte(fmt.Sprintf("%d", trade.Id))

	return k.SendMessage(topic, tradeID, tradeJSON)
}

func (k *KafkaProducerRepository) SendMatchTradeToKafkaSpecifyTopic(matchedTrades []domain.MatchingTrade, topic string) error {

	tradeJSON, err := json.Marshal(matchedTrades)
	if err != nil {
		return err
	}

	tradeID := []byte(fmt.Sprintf("%d", matchedTrades[0].Timestamp.Unix()))

	return k.SendMessage(topic, tradeID, tradeJSON)
}
