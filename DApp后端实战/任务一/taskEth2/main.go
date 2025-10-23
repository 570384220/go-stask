package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"taskEth2/store"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/11ae24000a504820bcb6ebc766923ac4")
	if err != nil {
		log.Fatal("client connect err : ", err)
	}
	// task 2
	keyStr, err := os.ReadFile("F:\\go_work_space\\sepoliaTestNetPrivateKey.txt")
	if err != nil {
		log.Fatal("read file err : ", err)
	}
	privateKey, err := crypto.HexToECDSA(string(keyStr))
	if err != nil {
		log.Fatal("private key err : ", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("get chain id err : ", err)
	}

	contractAddress := common.HexToAddress("0x24fFAE672839726dA606A2a61E31f43b9f964758")
	storeInstace, err := store.NewStore(contractAddress, client)
	if err != nil {
		log.Fatal("create store err : ", err)
	}

	var key, value [32]byte
	copy(key[:], []byte("hello"))
	copy(value[:], []byte("world"))

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal("new key err : ", err)
	}

	tx, err := storeInstace.SetItem(opts, key, value)
	if err != nil {
		log.Fatal("set item err : ", err)
	}
	fmt.Println("tx : ", tx.Hash().Hex())

	opts, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal("new key err : ", err)
	}
	returnValue, err := storeInstace.Items(&bind.CallOpts{
		Context: context.Background(),
	}, key)
	if err != nil {
		log.Fatal("get items err : ", err)
	}

	fmt.Printf("returnValue : %s\n", returnValue)

}
