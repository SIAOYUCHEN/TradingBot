package domain

import (
	"time"
)

type Trades struct {
	Id        int64          `gorm:"primaryKey"`        // 交易ID
	Market    TradeMarket    `gorm:"type:varchar(255)"` // 市场代码
	Price     float64        `gorm:"column:Price"`
	Amount    float64        `gorm:"column:Amount"`
	Direction TradeDirection `gorm:"type:varchar(255)"`
	Timestamp time.Time      `gorm:"type:timestamp;column:Timestamp"` // 时间戳记
}

type MatchingTrades struct {
	Id        int64       `gorm:"primaryKey"`
	AskId     int64       `gorm:"column:AskId"`
	BidId     int64       `gorm:"column:BidId"`
	Market    TradeMarket `gorm:"type:varchar(255);column:Market"`
	Price     float64     `gorm:"column:Price"`
	Amount    float64     `gorm:"column:Amount"`
	Timestamp time.Time   `gorm:"type:timestamp;column:Timestamp"` // 时间戳记
}

func (Trades) TableName() string {
	return "order"
}

func (MatchingTrades) TableName() string {
	return "transaction"
}
