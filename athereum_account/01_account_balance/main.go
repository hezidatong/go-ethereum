package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://34.80.53.52:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect success!")

	account := common.HexToAddress("0x35A21232691Da5B06A088F05004Fe07e643434F1")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(balance)
	fmt.Println("*****************")

	blockNumber := big.NewInt(204891)
	balance, err = client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	fmt.Println("++++++++++++++++++")

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
	pendingBalanceAt, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalanceAt)

}
