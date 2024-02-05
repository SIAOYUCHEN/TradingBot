package domain

import "fmt"

type TradeDirection int
type TradeMarket string

type Trade struct {
	Market    TradeMarket    `json:"market"`    // 市場代碼
	Price     float64        `json:"price"`     // 交易價格
	Amount    float64        `json:"amount"`    // 交易數量
	Direction TradeDirection `json:"direction"` // 交易方向，ask 表示賣出，bid 表示買入
}

const (
	Ask TradeDirection = iota // 賣出
	Bid                       // 買入
)

const (
	Taiwan    TradeMarket = "Taiwan"
	Japan     TradeMarket = "Japan"
	Usa       TradeMarket = "Usa"
	Singapore TradeMarket = "Singapore"
)

func (d TradeDirection) String() string {
	switch d {
	case Ask:
		return "0"
	case Bid:
		return "1"
	default:
		return fmt.Sprintf("TradeDirection(%d)", d)
	}
}
