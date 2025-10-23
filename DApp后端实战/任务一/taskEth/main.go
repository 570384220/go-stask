package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/11ae24000a504820bcb6ebc766923ac4")
	if err != nil {
		log.Fatal("client connect err : ", err)
	}

	number := big.NewInt(9469859)
	block, err := client.BlockByNumber(context.Background(), number)
	if err != nil {
		log.Fatal("get Block err :", err)
	}

	block.Time()

	// task 1
	log.Printf("block number : %d", block.Number())
	log.Printf("block hash : %s", block.Hash().Hex())
	log.Printf("block time : %s", time.Unix(int64(block.Header().Time), 0))
	log.Printf("block count : %d", block.Size())

	// task 2
	keyStr, err := os.ReadFile("F:\\go_work_space\\sepoliaTestNetPrivateKey.txt")
	if err != nil {
		log.Fatal("read file err : ", err)
	}
	privateKey, err := crypto.HexToECDSA(string(keyStr))
	if err != nil {
		log.Fatal("private key err : ", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	Nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("get nonce err : ", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("get gas price err : ", err)
	}

	value := big.NewInt(1000000000000000) // 0.001ETH

	toAddress := common.HexToAddress("0x15550f0D43492B67474f59f534A01af8966dce55")

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    Nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      3000000,
		GasPrice: gasPrice,
		Data:     []byte{},
	})

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("get chain id err : ", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal("sign tx err : ", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("send tx err : ", err)
	}

	fmt.Printf("tx hash : %s\n", signedTx.Hash().Hex())
}
