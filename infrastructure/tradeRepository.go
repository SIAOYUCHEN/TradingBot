package infrastructure

import (
	tradeInterface "TradingBot/domain/Interface"
	domain "TradingBot/domain/dto"
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type RedisTradeRepository struct {
	Client *redis.Client
}

func NewRedisTradeRepository(client *redis.Client) tradeInterface.TradeRedisRepository {
	return &RedisTradeRepository{
		Client: client,
	}
}

func (r *RedisTradeRepository) AddTrade(ctx context.Context, trade *domain.Trade, key string) error {

	tradeJSON, err := json.Marshal(trade)
	if err != nil {
		return err
	}

	if err := r.Client.RPush(ctx, key, tradeJSON).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RedisTradeRepository) GetTrade(ctx context.Context, key string) ([]*domain.Trade, error) {
	tradesJSON, err := r.Client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var trades []*domain.Trade
	for _, tradeJSON := range tradesJSON {
		var trade domain.Trade
		if err := json.Unmarshal([]byte(tradeJSON), &trade); err != nil {
			return nil, err
		}
		trades = append(trades, &trade)
	}

	return trades, nil
}

func (r *RedisTradeRepository) GetAllTrades(ctx context.Context) (map[string][]domain.Trade, error) {

	keys, err := r.Client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	tradesWithKeys := make(map[string][]domain.Trade)
	for _, key := range keys {
		tradesJSON, err := r.Client.LRange(ctx, key, 0, -1).Result()
		if err != nil {
			return nil, err
		}

		var tradesForKey []domain.Trade
		for _, tradeJSON := range tradesJSON {
			var trade domain.Trade
			err = json.Unmarshal([]byte(tradeJSON), &trade)
			if err != nil {
				return nil, err
			}
			tradesForKey = append(tradesForKey, trade)
		}
		tradesWithKeys[key] = tradesForKey
	}

	return tradesWithKeys, nil
}

func (r *RedisTradeRepository) ExistKey(ctx context.Context, key string) (int64, error) {
	return r.Client.Exists(ctx, key).Result()
}

func (r *RedisTradeRepository) DeleteTrade(ctx context.Context, key string) (int64, error) {
	return r.Client.Del(ctx, key).Result()
}
