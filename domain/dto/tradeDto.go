package domain

import (
	"time"
)

type TradeDirection string
type TradeMarket string

type Trade struct {
	Id        int64          `json:"id"`        // 交易ID
	Market    TradeMarket    `json:"market"`    // 市場代碼
	Price     float64        `json:"price"`     // 交易價格
	Amount    float64        `json:"amount"`    // 交易數量
	Direction TradeDirection `json:"direction"` // 交易方向，ask 表示賣出，bid 表示買入
	Timestamp time.Time      `json:"timestamp"` // 时间戳记
}

type MatchingTrade struct {
	AskId     int64       `json:"askid"`
	BidId     int64       `json:"bidid"`
	Market    TradeMarket `json:"market"`
	Price     float64     `json:"price"`
	Amount    float64     `json:"amount"`
	Timestamp time.Time   `json:"timestamp"`
}

const (
	Ask TradeDirection = "Ask" // 賣出
	Bid TradeDirection = "Bid" // 買入
)

const (
	Eth  TradeMarket = "ETH/USDT"
	Btc  TradeMarket = "BTC/USDT"
	Flow TradeMarket = "FLOW/USDT"
	Sol  TradeMarket = "SOL/USDT"
)

func (tm TradeMarket) IsValid() bool {
	switch tm {
	case Eth, Btc, Flow, Sol:
		return true
	default:
		return false
	}
}

func (td TradeDirection) IsValid() bool {
	switch td {
	case Ask, Bid:
		return true
	default:
		return false
	}
}
