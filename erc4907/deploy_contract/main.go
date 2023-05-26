package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	fmt.Println(fromAddress.Hex())
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
	name := "fangyan"
	symbol := "yan"

	address, tx, instance, err := erc4907.DeployErc4907(auth, client, name, symbol)
	if err != nil {
		log.Fatal(err)
	}

	//1*********************, nonce: 301
	//contract address: 0xd60F8A867a5625d97858DFF30aC3FCcFB4094E0D
	//transaction: 0xe22f5e2fe94220976ba2ea8cea1edea1f8c2a616368cbae28903e37ac0f1a8b4
	fmt.Println("contract address:", address)
	fmt.Println("transaction:", tx.Hash())

	_ = instance
}
