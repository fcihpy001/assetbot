package model

import "gorm.io/gorm"

// 钱包用户数据
type Trade struct {
	gorm.Model
	TxHash      string `gorm:"tx_hash;uniqueIndex:tx_type_idx;type:char(70)"`
	BlockNumber uint
	From        string
	To          string
	ETHAmount   string
	BlockTime   string
	Status      uint
}

type BlockInfo struct {
	gorm.Model
	BlockNumber int
}
