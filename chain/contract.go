package chain

import (
	"AssetBot/utils"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

func parseAirdropTransaction(tx *types.Transaction) {
	//获取交易凭据
	receipt, err := utils.GetEthClient().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	//遍历事件列表，找到感兴趣的事件
	for _, event := range receipt.Logs {

		if event.Address.Hex() == "" {

		}
	}
}
