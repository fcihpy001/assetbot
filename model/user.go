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
	UserId           string
	Name             string
	ex               string
	CreateTime       string
	AppManagerLevel  string
	GlobalRecvMsgOpt string
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
