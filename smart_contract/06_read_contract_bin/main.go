package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/febbb3f51b974458b5b3e1831a9fd11b")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x5d4c4227bB95B31cb458505783B98973d3646ff6")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode))
}
