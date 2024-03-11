package domain

import (
	dto "TradingBot/domain/dto"
)

type KafkaProducerRepository interface {
	SendTradeToKafkaSpecifyTopic(trade *dto.Trade, topic string) error

	SendMatchTradeToKafkaSpecifyTopic(matchedTrades []dto.MatchingTrade, topic string) error
}
