package main

import (
	"AssetBot/chain"
	"AssetBot/utils"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("hello")

	utils.InitConfig()

	//tt, _ := service.GetDB().Get("tt")
	//fmt.Println(tt)
	chain.ScanBlock()
}
