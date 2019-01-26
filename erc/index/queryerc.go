package main

import (
	"fmt"
	"log"

	"github.com/ShyftNetwork/go-empyrean/accounts/abi/bind"
	"github.com/ShyftNetwork/go-empyrean/common"
	"github.com/ShyftNetwork/go-empyrean/erc"
	"github.com/ShyftNetwork/go-empyrean/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0x021246cef13b38f3cd1cc558fe73c2ba7a5f02f0")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		fmt.Println("error here?")
		log.Fatal(err)
	}
	address := common.HexToAddress("0x021246cef13b38f3cd1cc558fe73c2ba7a5f02f0")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		fmt.Println("or here?")
		log.Fatal(err)
	}
	fmt.Println("balance:::::", bal)
	fmt.Println("token", tokenAddress.Hex())

	fmt.Println("instance::::", instance)

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
}
