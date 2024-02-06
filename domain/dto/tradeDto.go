package domain

type TradeDirection string
type TradeMarket string

type Trade struct {
	Market    TradeMarket    `json:"market"`    // 市場代碼
	Price     float64        `json:"price"`     // 交易價格
	Amount    float64        `json:"amount"`    // 交易數量
	Direction TradeDirection `json:"direction"` // 交易方向，ask 表示賣出，bid 表示買入
}

const (
	Ask TradeDirection = "Ask" // 賣出
	Bid TradeDirection = "Bid" // 買入
)

const (
	Eth  TradeMarket = "Eth"
	Btc  TradeMarket = "Btc"
	Flow TradeMarket = "Flow"
	Sol  TradeMarket = "Sol"
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
