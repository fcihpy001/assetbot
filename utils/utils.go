package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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
