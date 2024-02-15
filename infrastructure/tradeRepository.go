package infrastructure

import (
	tradeInterface "TradingBot/domain/Interface"
	domain "TradingBot/domain/dto"
	"context"
	"encoding/json"

	"time"

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

	// tradeJSON, err := json.Marshal(trade)
	// if err != nil {
	// 	return err
	// }

	// if err := r.Client.RPush(ctx, key, tradeJSON).Err(); err != nil {
	// 	return err
	// }

	// return nil

	score := float64(trade.Timestamp.UnixNano())
	member, err := json.Marshal(trade)
	if err != nil {
		return err
	}

	_, err = r.Client.ZAdd(ctx, key, &redis.Z{
		Score:  score,
		Member: member,
	}).Result()

	return err
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

func (r *RedisTradeRepository) MatchTrade(ctx context.Context, trade *domain.Trade, oppositeKey string, key string) error {

	oppositeTrades, err := r.Client.ZRangeByScore(ctx, oppositeKey, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  -1,
	}).Result()

	if err != nil {
		return err
	}

	var transactions []domain.MatchingTrade
	pipe := r.Client.TxPipeline()

	for _, oppTradeJSON := range oppositeTrades {
		var oppTrade domain.Trade
		err := json.Unmarshal([]byte(oppTradeJSON), &oppTrade)
		if err != nil {
			continue
		}

		if trade.Amount == 0 {
			break
		}

		if trade.Direction == "Bid" /*&& trade.Amount >= oppTrade.Amount*/ && trade.Price >= oppTrade.Price {

			transaction := domain.MatchingTrade{
				AskId:     trade.Id,
				BidId:     oppTrade.Id,
				Market:    trade.Market,
				Price:     oppTrade.Price,
				Amount:    oppTrade.Amount,
				Timestamp: time.Now(),
			}

			transactions = append(transactions, transaction)

			trade.Amount -= oppTrade.Amount
			oppTrade.Amount -= oppTrade.Amount

			if oppTrade.Amount <= 0 {
				pipe.ZRem(ctx, oppositeKey, oppTradeJSON)
			} else {
				updatedOppTradeJSON, _ := json.Marshal(oppTrade)
				pipe.ZAdd(ctx, oppositeKey, &redis.Z{
					Score:  float64(oppTrade.Timestamp.UnixNano()),
					Member: updatedOppTradeJSON,
				})
			}
		} else if trade.Direction == "Ask" /*&& trade.Amount >= oppTrade.Amount*/ && trade.Price <= oppTrade.Price {

			transaction := domain.MatchingTrade{
				AskId:     trade.Id,
				BidId:     oppTrade.Id,
				Market:    trade.Market,
				Price:     oppTrade.Price,
				Amount:    oppTrade.Amount,
				Timestamp: time.Now(),
			}

			transactions = append(transactions, transaction)

			trade.Amount -= oppTrade.Amount
			oppTrade.Amount -= oppTrade.Amount

			if oppTrade.Amount <= 0 {
				pipe.ZRem(ctx, oppositeKey, oppTradeJSON)
			} else {
				updatedOppTradeJSON, _ := json.Marshal(oppTrade)
				pipe.ZAdd(ctx, oppositeKey, &redis.Z{
					Score:  float64(oppTrade.Timestamp.UnixNano()),
					Member: updatedOppTradeJSON,
				})
			}
		}
	}

	if trade.Amount > 0 {
		newTradeJSON, _ := json.Marshal(trade)
		pipe.ZAdd(ctx, key, &redis.Z{
			Score:  float64(trade.Timestamp.UnixNano()),
			Member: newTradeJSON,
		})
	}

	for _, t := range transactions {
		transactionJSON, _ := json.Marshal(t)
		r.Client.RPush(ctx, "Transaction", transactionJSON)
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return err
}
