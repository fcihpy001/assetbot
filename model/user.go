package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type DidInfo struct {
	ID       uint   `gorm:"primaryKey"`
	Did      string `json:"did"`
	UUID     string `json:"uuid"`
	CoinType int
	Address  string
}

type User struct {
	gorm.Model
	TwitterId   string
	TwitterName string
	InviteId    uint
	HeadImage   string
	Pts         uint
	FFP         uint
	Version     uint
	Status      uint8
	Rank        uint
	BonkCount   uint
	ShieldCount uint
	WinCount    uint
	LossCount   uint
}

type AddressLoginParam struct {
	CoinType int    `form:"coin-type" binding:"required"`
	Address  string `form:"address"`
}

type HeaderData struct {
	Signature string `header:"signature"`
	DeviceId  string `header:"device-id" binding:"required"`
}

type DidClim struct {
	Did string `json:"did"`
	jwt.StandardClaims
}

type Bag struct {
	gorm.Model
	FFP         uint
	PTS         uint
	MintCount   uint
	BoxCount    uint
	Eths        uint
	Pets        uint
	InviteCount uint
	Shield      uint
}

type Food struct {
	gorm.Model
	Name  string
	Price uint
	Tod   uint
	Pts   uint
}

type Box struct {
	gorm.Model
	Type uint
	Name string
}

type Pet struct {
	gorm.Model
	Level uint8
	Body  uint
	Head  uint
	Leg   uint
	Mouse uint
	Eye   uint
}

type Game struct {
	gorm.Model
	Winner uint
	Loser  uint
	Income uint
	Loss   uint
}

// 交易记录
type Trade struct {
	gorm.Model
	Uid    uint
	FoodId Food
}

type Mint struct {
	gorm.Model
	UserID  User
	Pet     Pet
	EpochId uint
}

type SysConfig struct {
	gorm.Model
	//邀请几个奖励mint机会
	InviteCountRewardMint uint
	//每个epoch 时间
	EpochCoolTime uint
	//	默认mint 次数
	DefaultMintCount uint
	//	FFP 总数量
	TotalFFP uint
}

type FoodTrade struct {
	gorm.Model
	Uid         uint
	FoodId      uint
	Amount      uint
	Wallet      string
	TxHash      string `gorm:"tx_hash;uniqueIndex:tx_type_idx;type:char(66)"`
	BlockNumber int64
	From        string `json:"address"`
	To          string `json:"to"`
	ETHAmount   string `json:"eth_amount"`
	Eths        float64
	BlockTime   string `json:"block_time"`
	Status      uint64
}
