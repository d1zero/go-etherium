package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"math"
	"math/big"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := ethclient.DialContext(context.Background(), os.Getenv("URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	// get current block
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(block.Number())

	// get eth balance
	addr := common.HexToAddress("0x9e3DbA14ba046E00e17328fe653B23a13b62b53a")
	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	fmt.Println(new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow(10, 18))))

	// generate wallet
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err.Error())
	}
	privateData := crypto.FromECDSA(privateKey)
	publicData := crypto.FromECDSAPub(&privateKey.PublicKey)
	// private key
	fmt.Println(hexutil.Encode(privateData))
	// public key
	fmt.Println(hexutil.Encode(publicData))
	// address
	fmt.Println(crypto.PubkeyToAddress(privateKey.PublicKey).Hex())
}
