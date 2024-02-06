package domain

import (
	dto "TradingBot/domain/dto"
)

type DeleteTradeCommand struct {
	Market    dto.TradeMarket    `json:"market" validate:"required,validMarket"`
	Direction dto.TradeDirection `json:"direction" validate:"required,validDirection"`
}
