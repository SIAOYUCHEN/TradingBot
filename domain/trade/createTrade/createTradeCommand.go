package domain

import (
	dto "TradingBot/domain/dto"
)

type CreateTradeCommand struct {
	Market    dto.TradeMarket    `json:"market" validate:"required,validMarket"`
	Price     float64            `json:"price" validate:"required,gt=0"`
	Amount    float64            `json:"amount" validate:"required,gt=0"`
	Direction dto.TradeDirection `json:"direction" validate:"required,validDirection"`
}
