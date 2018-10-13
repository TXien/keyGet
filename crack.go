

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
	"os"
	"time"
	//"reflect"
)

func main() {
        mainRun()
}

func mainRun(){
        //client, err := ethclient.Dial("http://192.168.51.203:9999")//"https://mainnet.infura.io")
	client, err := ethclient.Dial("https://mainnet.infura.io")
        if err != nil {
                mainRun()//log.Fatal(err)
        }
        for i:=0;i<=10000;i++{
                go run(client,i)
        }
        time.Sleep(10086500 * time.Second)
}

func run(client *ethclient.Client, km int){
	for i:=0;i!=-1;i++{
		//if(i%1000==0 &&km == 0){
			fmt.Println(i)
		//}
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
	//os.Create(addressTarget+"___"+hex.EncodeToString(priv.D.Bytes()))
	if(balance != "0" ){
		os.Create(addressTarget+"___"+hex.EncodeToString(priv.D.Bytes()))
	        fmt.Printf("%x\n", priv.D.Bytes())
	        fmt.Println(addressTarget)
	        fmt.Println(getBalance(client, addressTarget))
		/*f, err := */os.Create(hex.EncodeToString(priv.D.Bytes()))
	}
}

func getBalance(client *ethclient.Client, address string)(string){
        header, err := client.BalanceAt(context.Background(),common.HexToAddress(address), nil)
        if err != nil {
                //log.Fatal(err)
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


