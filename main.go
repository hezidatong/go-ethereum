package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum/erc4907"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://192.168.0.164:8545")
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xeec720FfB4b7106922206f21087b4b29C7201cbd")
	at, _ := client.NonceAt(context.Background(), toAddress, nil)
	fmt.Println("at::", at)

	nonce, err := client.PendingNonceAt(context.Background(), toAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2*********************, nonce::", nonce)
	// return

	_, tx, instance, auth, err := deployContract(client)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil || receipt.Status == 0 {
			i++
			fmt.Println(i)
			continue
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

		tokenId := big.NewInt(1)
		toAddress := common.HexToAddress("0xeec720FfB4b7106922206f21087b4b29C7201cbd")

		nonce, err := client.PendingNonceAt(context.Background(), toAddress)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("2*********************, nonce::", nonce)

	label:
		tx2, err := instance.Mint(&bind.TransactOpts{
			From:      toAddress,
			Nonce:     big.NewInt(int64(nonce)),
			Signer:    SignerFn,
			Value:     auth.Value,
			GasPrice:  auth.GasPrice,
			GasFeeCap: nil,
			GasTipCap: nil,
			GasLimit:  auth.GasLimit,
			Context:   context.Background(),
			NoSend:    false,
		}, tokenId, toAddress)
		if err != nil {
			log.Fatal("Mint err:", err)
		}

		receipt, err = client.TransactionReceipt(context.Background(), tx2.Hash())
		if err != nil || receipt.Status == 0 {
			i++
			fmt.Println(i)
			goto label
		}

		return
	}

}

func SignerFn(addr common.Address, tx *types.Transaction) (*types.Transaction, error) { return tx, nil }

func deployContract(client *ethclient.Client) (common.Address, *types.Transaction, *erc4907.Erc4907, *bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA("0c59ab815bc011b396656c24a98cd6513692aa27b4740740c25a2745de869da1")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Println("1*********************, nonce:", nonce)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//transactor := bind.NewKeyedTransactor(privateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	auth.GasPrice = gasPrice
	name := "xubo"
	symbol := "bo"

	address, tx, instance, err := erc4907.DeployErc4907(auth, client, name, symbol)
	if err != nil {
		log.Fatal(err)
	}
	return address, tx, instance, auth, err
}
