package utils

import (
	"AssetBot/model"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"math/big"
	"os"
	"time"
)

var (
	Config      model.Config
	ethClient   *ethclient.Client
	ethWSClient *ethclient.Client
)

func InitConfig() {
	yamlFile, err := os.ReadFile("./config.yaml")
	if os.Getenv("DEBUG") == "1" {
		yamlFile, err = os.ReadFile("./config_dev.yaml")
	}
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return
	}
	log.Println("配置文件读取成功")

}

func GetABI(abiJson string) abi.ABI {
	file, err := os.ReadFile(abiJson)
	if err != nil {
		log.Panicln("文件读取失败")
	}
	wrapABI, err := abi.JSON(bytes.NewReader(file))
	return wrapABI
}

func GetEthClient() *ethclient.Client {
	if ethClient == nil {
		// 加载环境变量从 .env 文件
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		log.Println("env 文件读取功")
		log.Println(Config.RPC)
		dial, err := ethclient.Dial(Config.RPC)
		if err != nil {
			log.Fatalf("连接以太坊节点失败：%v", err)
		}
		ethClient = dial
		fmt.Println("eth_client节点初始化成功")
	}
	return ethClient
}

func GetEthWSClient() *ethclient.Client {
	if ethWSClient == nil {
		dial, err := ethclient.Dial("https://bsc-testnet.public.blastapi.io")
		if err != nil {
			log.Fatalf("连接以太坊节点失败：%v", err)
		}
		ethWSClient = dial
		fmt.Println("eth_client节点初始化成功")
	}
	return ethWSClient
}

func BigIntToUint(i *big.Int) uint {
	// 检查是否为负数
	if i.Sign() == -1 {
		return 0
	}

	// 将 big.Int 转换为 *big.Float
	f := new(big.Float).SetInt(i)

	// 将 *big.Float 转换为 uint
	uintVal, _ := f.Uint64()

	return uint(uintVal)
}

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//func Time(timeInterval uint64) time.Time {
//
//}

func HexSum(hex1 string, hex2 string) string {
	if len(hex1) == 0 {
		return hex2
	}
	if len(hex2) == 0 {
		return hex1
	}
	value1 := new(big.Int)
	v1, _ := value1.SetString(hex1, 16)

	value2 := new(big.Int)
	v2, _ := value2.SetString(hex2, 16)

	sum := new(big.Int)
	sum.Add(v1, v2)
	return fmt.Sprintf("%d", sum)
}

func ConverTime(iso8601 string) string {
	//这个必须加UTC参数，不加本地测试和服务器上不一致
	t, err := time.ParseInLocation(time.RFC3339, iso8601, time.UTC)
	if err != nil {
		return iso8601
	}
	//设置东八区
	cz := time.FixedZone("CST", 8*3600)
	return t.In(cz).Format("2006-01-02 15:04:05")
}

func FormatAmount(data []uint8) string {
	amountInt := new(big.Int)
	amountInt.SetBytes(data)
	amountValue := fmt.Sprintf("%d", amountInt)
	return amountValue
}

func GetUsdtAmount(data []uint8) string {
	if len(data) < 18 {
		return ""
	}
	b := new(big.Int)
	value := b.SetBytes(data[:])
	amount := fmt.Sprintf("%d", value)
	length := len(amount)
	return amount[:length-18]
}
