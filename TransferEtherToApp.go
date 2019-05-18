package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {	

	// client, err := ethclient.Dial("http://127.0.0.1:7545")
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("091D55FC208023DC94F95E129F8608DDF36BC7BEC782B9A47B1B7A50495F2E45")
	
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
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(6000000) // in units
	auth.GasPrice = gasPrice
	address := common.HexToAddress("0x4a59C5A565E7Ff16F66D9690D3Cbcb47DC0DAe63")
	instance, err := NewCryptoContract(address, client)

	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(105) 
	if err != nil {
		log.Fatal(err)
	}
	
	toAddress := common.HexToAddress("0xE33f4C2306eFE9BF66a64A3c42408d2Fe1Cb890f")

	tx, err := instance.Transfer(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 2381623,
		Value:    big.NewInt(305),
	}, toAddress, value)
	
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("tx_transfer sent: %s", tx.Hash().Hex()) 
	result, err := instance.GetBalance(&bind.CallOpts{})
	
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("Total Balance of %s : %s", result) 
}
