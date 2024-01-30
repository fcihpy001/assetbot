package model

import "gorm.io/gorm"

// 钱包用户数据
type ChainTrade struct {
	gorm.Model
	TxHash      string `gorm:"tx_hash;uniqueIndex:tx_type_idx;type:char(66)"`
	BlockNumber int64
	From        string  `json:"address"`
	To          string  `json:"to"`
	ETHAmount   string  `json:"eth_amount"`
	Eth         float64 `json:"eth"`
	BlockTime   string  `json:"block_time"`
	Status      uint64
}
