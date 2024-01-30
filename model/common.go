package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorInfo struct {
}
type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ErrorResp(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, &ApiResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func SuccessResp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ApiResponse{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

type HttpData struct {
	Result Result
}

type Result struct {
	TransactionHash string       `json:"transaction_hash"`
	Status          string       `json:"status"`
	Function        string       `json:"function"`
	Logs            []EthersInfo `json:"logs"`
	BlockNumber     string       `json:"block_number"`
	BlockTimestamp  string       `json:"block_timestamp"`
	From            string       `json:"from"`
	To              string       `json:"to"`
	ContractAddress string       `json:"effective_contract_address"`
}

type EthersInfo struct {
	Data            EthersData
	Event           string
	ContractType    string
	ContractAddress string
}

type EthersData struct {
	From   string
	To     string
	Amount string
}
