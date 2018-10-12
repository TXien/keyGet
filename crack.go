

package main

import (
	"context"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"log"
	//"math/big"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
	//"reflect"
)

func main() {
        client, err := ethclient.Dial("http://192.168.51.203:9999")//"https://mainnet.infura.io")
        if err != nil {
                log.Fatal(err)
        }
	for i:=0;i<=50;i++{
		go run(client)
	}
	time.Sleep(10086500 * time.Second)
}

func run(client *ethclient.Client){
	for i:=0;i!=1;i++{
		if(i%100000==0){
			fmt.Println(i)
		}
                verifyAccount(client)
        }
}

func verifyAccount(client *ethclient.Client){
        priv, _ := crypto.GenerateKey()
        pub := priv.PublicKey
        ecdsaPubBytes := elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)
        addressBytes := crypto.Keccak256(ecdsaPubBytes[1:])[12:]
        addressTarget := hex.EncodeToString(addressBytes)
	balance := getBalance(client, addressTarget)
	//fmt.Println(addressTarget)
	//fmt.Println(balance)
	//fmt.Printf("%x\n", priv.D.Bytes())
	if(balance != "0" ){
	        fmt.Printf("%x\n", priv.D.Bytes())
	        fmt.Println(addressTarget)
	        fmt.Println(getBalance(client, addressTarget))
	}
}

func getBalance(client *ethclient.Client, address string)(string){
        header, err := client.BalanceAt(context.Background(),common.HexToAddress(address), nil)
        if err != nil {
                log.Fatal(err)
        }
        return header.String()
}

func blockNumber(client *ethclient.Client)(string){
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return header.Number.String()
}
