package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum/smart_contract/02_deploy/storage"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/febbb3f51b974458b5b3e1831a9fd11b")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x5d4c4227bB95B31cb458505783B98973d3646ff6")
	instance, err := storage.NewStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
