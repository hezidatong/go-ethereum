package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum/erc4907"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://192.168.0.164:8545")
	if err != nil {
		log.Fatal("client err: ", err)
	}

	// 合约地址
	address := common.HexToAddress("0x946417F172F710dFc2A999529BDeB6d7CE281995")

	// mint nft toAddress
	toAddress := common.HexToAddress("0xeec720FfB4b7106922206f21087b4b29C7201cbd")

	instance, err := erc4907.NewErc4907(address, client)
	if err != nil {
		log.Fatal("instance err: ", err)
	}

	// 查询合约
	name1, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal("name err:", err)
	}
	fmt.Println("name:", name1)

	symbol1, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("symbol:", symbol1)

	at, _ := client.NonceAt(context.Background(), toAddress, nil)
	fmt.Println("at::", at)

	nonce, err := client.PendingNonceAt(context.Background(), toAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2*********************, nonce::", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("0c59ab815bc011b396656c24a98cd6513692aa27b4740740c25a2745de869da1")
	if err != nil {
		log.Fatal(err)
	}

	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	auth.GasPrice = gasPrice

	tokenId := big.NewInt(1)
	tx, err := instance.Mint(auth, tokenId, toAddress)
	if err != nil {
		log.Fatal("Mint err:", err)
	}

	fmt.Println("tx:", tx.Hash())

	i := 0
	for {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil || receipt.Status == 0 {
			i++
			fmt.Println(i)
			continue
		}

		tokenURI, err := instance.TokenURI(&bind.CallOpts{}, tokenId)
		if err != nil {
			log.Fatal("TokenURI err:", err)
		}
		fmt.Println("tokenURI:", tokenURI)
		return
	}

}
