package token

import (
	"encoding/hex"
	"fmt"
	"strings"

	"context"
	"log"

	"math/big"

	"github.com/ShyftNetwork/go-empyrean"
	"github.com/ShyftNetwork/go-empyrean/accounts/abi"
	"github.com/ShyftNetwork/go-empyrean/accounts/abi/bind"
	"github.com/ShyftNetwork/go-empyrean/common"
	"github.com/ShyftNetwork/go-empyrean/core/types"
	"github.com/ShyftNetwork/go-empyrean/crypto"
	"github.com/ShyftNetwork/go-empyrean/erc"
	"github.com/ShyftNetwork/go-empyrean/ethclient"
)

func WriteTokenTransfers(data []byte, addr string) {
	var count = 0
	erc20Standard := [6]string{"dd62ed3e", "095ea7b3", "70a08231", "18160ddd", "a9059cbb", "23b872dd"}
	str := hex.EncodeToString(data)
	for _, i := range erc20Standard {
		if strings.Contains(str, i) {
			count++
		}
	}
	fmt.Println(addr)
	isERC := false
	if count == 6 {
		isERC = true
		go LogTokenTransferData(addr)
	}
	fmt.Println(isERC)
}

func LogTokenTransferData(addr string) {
	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		fmt.Println(err)
	}

	tokenAddress := common.HexToAddress(addr)
	instance, err := erc.NewErc(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}
	total, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}
	address := common.HexToAddress(addr)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		fmt.Println(err)
	}

	EventSubscription(tokenAddress, client)

	fmt.Println("TOTAL SUPPLY:::::", total)
	fmt.Println("BALANCE     :::::", bal)
	fmt.Println("CONTRACT ADDR::::", tokenAddress.Hex())
	fmt.Println("TOKEN NAME   ::::", name)
	fmt.Println("TOKEN SYMBOL ::::", symbol)
	fmt.Println("TOKEN DECIMAL :::", decimals)
}

func EventSubscription(addr common.Address, client *ethclient.Client) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("EVENT SUBSCRIPTION", vLog) // pointer to event log
		}
	}
}

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func LoggingERC20Events(addr string, client *ethclient.Client) {
	contractAddress := common.HexToAddress(addr)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(360),
		ToBlock:   big.NewInt(700),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(erc.ErcABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		fmt.Printf("\n\n")
	}
}
