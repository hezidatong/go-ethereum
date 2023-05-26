package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum/smart_contract/07_donft/contracts"
	"math/big"
	"strconv"
)

const (
	Irishub    = "6666"
	Uptick     = "6667"
	IrisHost   = "http://34.80.53.52"
	UptickHost = "http://34.80.181.4"
	APIPort    = ":1317"
	Route      = "/erc721_bridge/converter/v1/token_trace/"
	ChainParam = "erc-721/channel-0/"
	ERCPort    = ":8545"
)

func main() {
	// iris链 host = http://34.80.53.52:8545, donft_contract_addr = 0x7e4abfba5aa9d3484c162bc882b4150f52742f2a
	// uptick链 host = http://34.80.181.4:8545, donft_contract_addr = 0xb514682a62ff1ea9f45d297d53b13786a76a6ba3
	queryHost := IrisHost

	//"http://34.80.53.52:8545"
	client, err := ethclient.Dial(queryHost + ERCPort)
	if err != nil {
		fmt.Println("err1:", err)
	}

	address := common.HexToAddress("0x7e4abfba5aa9d3484c162bc882b4150f52742f2a")
	instance, err := contracts.NewContracts(address, client)
	if err != nil {
		fmt.Println("err2:", err)
	}

	//bb := new(big.Int)
	//tokenID, _ := bb.SetString("27", 10)

	atoi, _ := strconv.Atoi("27")
	tokenID := big.NewInt(int64(atoi))

	ownerOf, err := instance.OwnerOf(&bind.CallOpts{}, tokenID)
	if err != nil {
		fmt.Println("err3:", err)
	}
	fmt.Println(ownerOf.Hash())
	fmt.Println(ownerOf.Hex())
	fmt.Println(ownerOf.String()[2:])
}
