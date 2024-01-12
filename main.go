package main

import (
	"AssetBot/chain"
	"AssetBot/utils"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	utils.InitConfig()

	chain.ScanChaiTask()
}
