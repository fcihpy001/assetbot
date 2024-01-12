package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

func HashToInt(hash common.Hash) uint {
	bigInt := new(big.Int)
	bigInt.SetBytes(hash[:])
	return BigIntToUint(bigInt)
}

func HashToUint64(hash common.Hash) uint64 {
	bigInt := new(big.Int)
	bigInt.SetBytes(hash[:])
	return bigInt.Uint64()
}

func BlockTime(timestamp uint64) string {
	// 你的Unix时间戳，以秒为单位
	timestampInterval := int64(timestamp)

	// 创建一个Time类型的值，使用Unix函数将时间戳转换为Time
	t := time.Unix(timestampInterval, 0)

	// 设置时区为东八区（中国标准时间）
	cst := time.FixedZone("CST", 8*60*60)

	// 使用In函数将时间转换为指定时区的时间
	cstTime := t.In(cst)
	// 格式化为日期时间字符串
	formattedTime := cstTime.Format("2006-01-02 15:04:05")
	return formattedTime
}
