package chain

import (
	"AssetBot/model"
	"AssetBot/service"
	"AssetBot/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"
)

var (
	startBlock int64 = 10309326
	isRunning  bool  = false
	blockTime  string
)

func Task() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				ScanBlock()
			}
		}
	}()
	<-make(chan interface{})
}

func ScanBlock() {
	if isRunning {
		return
	}
	isRunning = true

	//获取当前区块高度
	header, err := utils.GetEthClient().HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Printf("获取区块高度有误:%v", err)
		return
	}
	log.Println("获取最新高度:", header.Number.Int64())

	startBlock = int64(10309326)
	log.Println("开始执行扫描程序...", startBlock)
	//for i := startBlock; i <= header.Number.Int64(); i++ {
	//	getBlockInfo(i)
	//}
	blocks := [...]int64{
		10309326, 10309263,
	}
	for _, i := range blocks {
		getBlockInfo(i)
	}
	isRunning = false
}

func getBlockInfo(blockNumber int64) {
	block, err := utils.GetEthClient().BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Println("获取区块信息出错:", err)
		return
	}
	log.Println("正在扫描区块：", blockNumber)
	blockTime = utils.BlockTime(block.Time())

	chainID, err := utils.GetEthClient().NetworkID(context.Background())

	for _, tx := range block.Transactions() {
		receipt, _ := utils.GetEthClient().TransactionReceipt(context.Background(), tx.Hash())

		//保存交易的发送方地址
		if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
			fmt.Println("from-", from.Hex())
			trade := model.Trade{
				TxHash:      tx.Hash().Hex(),
				BlockNumber: blockNumber,
				From:        from.Hex(),
				To:          tx.To().Hex(),
				ETHAmount:   tx.Value().String(),
				Status:      receipt.Status,
			}
			service.GetDB().Save(trade)
		}
	}

}

//func Subscriber() {
//	contractAddress := common.HexToAddress("0x1643E812aE58766192Cf7D2Cf9567dF2C37e9B7F")
//	query := ethereum.FilterQuery{
//		FromBlock: big.NewInt(2394201),
//		ToBlock:   big.NewInt(2394201),
//		Addresses: []common.Address{
//			contractAddress,
//		},
//	}
//
//	logs, err := utils.GetEthClient().FilterLogs(context.Background(), query)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, vLog := range logs {
//		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
//		fmt.Println(vLog.BlockNumber)     // 2394201
//		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
//
//		event := struct {
//			Key   [32]byte
//			Value [32]byte
//		}{}
//		//err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println(string(event.Key[:]))   // foo
//		fmt.Println(string(event.Value[:])) // bar
//
//		var topics [4]string
//		for i := range vLog.Topics {
//			topics[i] = vLog.Topics[i].Hex()
//		}
//
//		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
//	}
//
//	eventSignature := []byte("ItemSet(bytes32,bytes32)")
//	hash := crypto.Keccak256Hash(eventSignature)
//	fmt.Println(hash.Hex())
//
//}

//func GetBlocks() {
//
//	// 以太坊节点的RPC URL
//	//rpcURL := "https://bsc-testnet.public.blastapi.io"
//
//	contractAddr := utils.CouponAddress()
//	//contractABI := utils.GetABI("./ABI/Coupon.json")
//	rpcURL := "https://bsc-testnet.public.blastapi.io"
//	cl := &rpc.HTTP{
//		BaseURL: rpcURL,
//	}
//	// 查询与合约地址相关的所有交易
//	logs, err := cl.GetLogs(rpc.FilterParams{
//		Address: contractAddress,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 提取每个交易所在的区块号
//	for _, log := range logs {
//		fmt.Printf("Transaction Hash: %s, Block Number: %d\n", log.TxHash, log.BlockNumber)
//	}
//}
