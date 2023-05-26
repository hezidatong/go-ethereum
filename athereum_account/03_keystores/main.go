package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

func importKs() {
	file := "./tmp/UTC--2023-05-23T06-41-57.338365600Z--d2f812b05203acbdcea429a05521d4078696b562"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("**********")
		log.Fatal(err)
	}

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		fmt.Println("++++++++++++")
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

}

func main() {
	//createKs()
	importKs()
}
