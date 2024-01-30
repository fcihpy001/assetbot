package utils

import (
	"AssetBot/model"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
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

func Wei2ether(wei *big.Int) float64 {
	flotValue := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(flotValue, big.NewFloat(math.Pow10(18)))
	ethFlot, _ := ethValue.Float64()
	return ethFlot
}

func HttpRequest(tx string) *model.HttpData {
	requestUrl := fmt.Sprintf("https://api.facet.org/transactions/%s", tx)
	fmt.Println("本次请求url", requestUrl)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, requestUrl, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("请求错误：", err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var data model.HttpData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("解析 JSON 时出错:", err)
		return nil
	}
	return &data
}
