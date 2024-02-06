package domain

import (
	dto "TradingBot/domain/dto"
)

type GetAllTradeResponse struct {
	TradesWithKeys map[string][]dto.Trade `json:"trades"`
}
