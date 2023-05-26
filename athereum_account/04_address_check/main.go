package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("is valid: %v\n", re.MatchString("0x35A21232691Da5B06A088F05004Fe07e643434F1"))
	fmt.Printf("is valid: %v\n", re.MatchString("0xZ23b5d4c32345ced77393b3530b1eed0f346429d"))
	fmt.Println(len("0x323b5d4c32345ced77393b3530b1eed0f346429d"))

	client, err := ethclient.Dial("http://34.80.53.52:8545")
	if err != nil {
		log.Fatal(err)
	}

	// "0x35A21232691Da5B06A088F05004Fe07e643434F1"
	address := common.HexToAddress("0x7e4abfba5aa9d3484c162bc882b4150f52742f2a")
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract)

}
