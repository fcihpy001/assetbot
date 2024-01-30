package main

import (
	"AssetBot/utils"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	utils.InitConfig()
	//
	//chain.ScanChaiTask()

	//chain.GetBlockInfo(19097601)

	data := utils.HttpRequest("0x0f3a4b71f0935f26b11a0f5666ab4c6aea2eb9e9f6a2300d375441b7c40ef237")
	fmt.Printf("info:%v\n", data)
	for _, info := range data.Result.Logs {
		if info.Event == "Transfer" {
			fmt.Println("event", info.Event, "type", info.ContractType, "addr:", info.ContractAddress)
			ethersData := info.Data
			println("from:", ethersData.From, "to:", ethersData.To, "amount", ethersData.Amount)
		}
	}
}
