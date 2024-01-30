package chain

import (
	"AssetBot/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"
)

var (
	startBlock int64 = 0
	isRunning  bool  = false
	blockTime  string
)

func ScanChaiTask() {
	//从数据库中读取区块号
	var block model.BlockInfo
	service.GetDB().First(&block)
	startBlock = int64(block.BlockNumber)

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

// 0xc9fd9196a7612fA09309c624e272DbcF310881Ac
// 0xE813FCE98F43abC18Dc4befFeACa2d8B7c47521c
// 扫链方式
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
	if startBlock == 0 {
		startBlock = header.Number.Int64()
	}

	log.Println("开始执行扫描程序...", startBlock)
<<<<<<< HEAD
	for i := startBlock; i <= header.Number.Int64(); i++ {
		GetBlockInfo(i)
=======
	for i := startBlock; i <= startBlock+2; i++ {
		getBlockInfo(i)
>>>>>>> 5070448225c8f1a4a8f6811a48979cfd748c8cf1
	}

	//将当前区块高度写入文件
	var block model.BlockInfo
	service.GetDB().First(&block)
	if block.ID != 0 {
		result := service.GetDB().Model(&model.BlockInfo{}).Where("block_number", startBlock).Update("block_number", header.Number.Int64())
		if result.Error != nil {
			fmt.Println("写入失败", result)
		}
	} else {
		block.BlockNumber = int(startBlock)
		service.GetDB().Create(&block)
	}

	isRunning = false
}

func GetBlockInfo(blockNumber int64) {
	block, err := utils.GetEthClient().BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Println("获取区块信息出错:", err)
		return
	}
	log.Println("正在扫描区块：", blockNumber)
	blockTime = utils.BlockTime(block.Time())

	chainID, err := utils.GetEthClient().NetworkID(context.Background())

	for i, tx := range block.Transactions() {
		receipt, _ := utils.GetEthClient().TransactionReceipt(context.Background(), tx.Hash())
		//fmt.Println("status:", receipt.Status, "hash:", tx.Hash().Hex(), "-index", i)
		//保存交易的发送方地址
<<<<<<< HEAD
		from, err := types.Sender(types.NewLondonSigner(chainID), tx)
		//err == nil {
		//fmt.Println("from:", from, "to:", tx.To().Hex())
		//if tx.To().Hex() == "0x00000000000000000000000000000000000FacE7" {
		//fmt.Println("eths", utils.Wei2ether(tx.Value()))
		//fmt.Println("from-", from.Hex())
		//fmt.Println("hash", tx.Hash())
		//trade := model.ChainTrade{
		//	TxHash:      tx.Hash().Hex(),
		//	BlockNumber: blockNumber,
		//	From:        from.Hex(),
		//	To:          tx.To().Hex(),
		//	ETHAmount:   tx.Value().String(),
		//	Status:      receipt.Status,
		//	Eth:         utils.Wei2ether(tx.Value()),
		//	BlockTime:   blockTime,
		//}

		//if err := service.GetDB().Create(&trade).Error; err != nil {
		//	fmt.Println("入库失败", err)
		//}
		//	fmt.Println("插入记录成功:daddd")
		//}
		if err != nil {
			fmt.Println("errorwfqf:", err)
		} else {
			//fmt.Println("from:", from, "to:", tx.To().Hex())
			if from.Hex() == "0xE813FCE98F43abC18Dc4befFeACa2d8B7c47521c" && tx.To().Hex() == "0x00000000000000000000000000000000000FacE7" {

				fmt.Println("hash", tx.Hash().Hex())
				fmt.Println("face token", i)
				fmt.Println("recie", receipt.Status)
				//== "0x00000000000000000000000000000000000FacE7" &&
				fmt.Println("to", tx.To().Hex(), "from-", from.Hex())
			}

=======
		if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
			fmt.Println("from-", from.Hex())
			trade := model.Trade{
				TxHash:      tx.Hash().Hex(),
				BlockNumber: uint(blockNumber),
				From:        from.Hex(),
				To:          tx.To().Hex(),
				ETHAmount:   tx.Value().String(),
				BlockTime:   utils.BlockTime(block.Time()),
				Status:      uint(receipt.Status),
			}
			err := service.GetDB().Create(&trade).Error
			if err != nil {
				println("入库失败:", err)
			}
>>>>>>> 5070448225c8f1a4a8f6811a48979cfd748c8cf1
		}
	}
	fmt.Println("本block扫描完成")
}

//监听方式
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
