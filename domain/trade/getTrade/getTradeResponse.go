package domain

import (
	dto "TradingBot/domain/dto"
)

type GetTradeResponse struct {
	Trades []dto.Trade `json:"trades"`
}
