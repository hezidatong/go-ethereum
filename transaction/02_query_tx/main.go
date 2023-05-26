package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/febbb3f51b974458b5b3e1831a9fd11b")
	if err != nil {
		log.Fatal(err)
	}

	//blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		//fmt.Println(tx.Hash().Hex())
		//fmt.Println(tx.Value().String())
		//fmt.Println(tx.Gas())
		//fmt.Println(tx.GasPrice().Uint64())
		//fmt.Println(tx.Nonce())
		//fmt.Println(tx.Data())
		//fmt.Println(tx.To().Hex())

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(chainID)
		//break

		baseFee := big.NewInt(4000000)
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), baseFee); err == nil {
			fmt.Println(msg.From().Hex())
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
	}

	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex())
	}

	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending)

}
