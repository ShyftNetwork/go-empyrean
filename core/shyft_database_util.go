package core

import (
	"math/big"
	"strconv"
	"time"

	"database/sql"
	"flag"
	"fmt"
	"strings"

	"github.com/ShyftNetwork/go-empyrean/common"
	Rewards "github.com/ShyftNetwork/go-empyrean/consensus/ethash"
	"github.com/ShyftNetwork/go-empyrean/core/sTypes"
	"github.com/ShyftNetwork/go-empyrean/core/types"
	"github.com/ShyftNetwork/go-empyrean/ethdb"
	"github.com/ShyftNetwork/go-empyrean/log"
	"github.com/ShyftNetwork/go-empyrean/track"
)

var InternalTracker track.InternalTracker

// Sets up shyft tracer struct type
func SetInternalTracker(st track.InternalTracker) {
	InternalTracker = st
}

func WriteShyftGen(db ethdb.SDatabase, gen *Genesis, block *types.Block) {
	for k, v := range gen.Alloc {
		_, _, err := db.AccountExists(k.String())
		switch {
		case err == sql.ErrNoRows:
			var toAddr *common.Address
			var data []byte
			var gasPrice uint64
			var cost string
			//Initializing proper types for tx struct
			toAddr = &k
			cost = "0"
			gasPrice = 0
			//Appending GENESIS to address stored as txHash and From Addr
			Genesis := []string{"GENESIS_", k.String()}
			GENESIS := "GENESIS"
			txHash := strings.Join(Genesis, k.String())
			//Create the accountNonce, set to 1 (1 incoming tx), format type
			accountNonce := v.Nonce
			i, err := strconv.ParseInt(block.Time().String(), 10, 64)
			if err != nil {
				panic(err)
			}
			age := time.Unix(i, 0)
			txData := stypes.ShyftTxEntryPretty{
				TxHash:      txHash,
				From:        GENESIS,
				To:          toAddr.String(),
				BlockHash:   block.Header().Hash().Hex(),
				BlockNumber: block.Header().Number.String(),
				Amount:      v.Balance.String(),
				Cost:        cost,
				GasPrice:    gasPrice,
				GasLimit:    block.GasLimit(),
				Gas:         block.GasUsed(),
				Nonce:       accountNonce,
				Age:         age,
				Data:        data,
				Status:      "SUCCESS",
				IsContract:  false,
			}
			//Create account and store tx
			db.InsertTx(txData)

		default:
			log.Info("Found Genesis Block")
		}
	}
}

// WriteShyftBlockZero writes block 0 to postgres db
func WriteShyftBlockZero(db ethdb.SDatabase, block *types.Block, gen *Genesis) error {

	i, err := strconv.ParseInt(block.Time().String(), 10, 64)
	if err != nil {
		panic(err)
	}
	age := time.Unix(i, 0)

	blockData := stypes.SBlock{
		Hash:       block.Header().Hash().Hex(),
		Coinbase:   block.Header().Coinbase.String(),
		Number:     block.Header().Number.String(),
		GasUsed:    block.Header().GasUsed,
		GasLimit:   block.Header().GasLimit,
		TxCount:    len(gen.Alloc),
		UncleCount: len(block.Uncles()),
		Age:        age,
		ParentHash: block.ParentHash().String(),
		UncleHash:  block.UncleHash().String(),
		Difficulty: block.Difficulty().String(),
		Size:       block.Size().String(),
		Nonce:      block.Nonce(),
		Rewards:    "0",
	}
	exist := db.BlockExists(blockData.Hash)
	if !exist {
		db.InsertBlock(blockData)
		log.Info("Block zero written to DB")
	}
	return nil
}

//SWriteBlock writes to block info to sql db
func SWriteBlock(db ethdb.SDatabase, block *types.Block, receipts []*types.Receipt) error {
	//Get miner rewards
	rewards := swriteMinerRewards(db, block)
	//Format block time to be stored
	i, err := strconv.ParseInt(block.Time().String(), 10, 64)
	if err != nil {
		panic(err)
	}
	age := time.Unix(i, 0)

	blockData := stypes.SBlock{
		Hash:       block.Header().Hash().Hex(),
		Coinbase:   block.Header().Coinbase.String(),
		Number:     block.Header().Number.String(),
		GasUsed:    block.Header().GasUsed,
		GasLimit:   block.Header().GasLimit,
		TxCount:    block.Transactions().Len(),
		UncleCount: len(block.Uncles()),
		ParentHash: block.ParentHash().String(),
		UncleHash:  block.UncleHash().String(),
		Difficulty: block.Difficulty().String(),
		Size:       block.Size().String(),
		Nonce:      block.Nonce(),
		Rewards:    rewards,
		Age:        age,
	}

	//Inserts block data into DB
	db.InsertBlock(blockData)

	if block.Transactions().Len() > 0 {
		for _, tx := range block.Transactions() {
			// SHYFT NOTE: Some contract tests have an invalid value - so if this is the case
			// we dont write the transaction to the Database
			if tx.From() != nil {
				swriteTransactions(db, tx, block.Header().Hash(), blockData.Number, receipts, age, blockData.GasLimit)
			}
		}
	}
	return nil
}

//swriteTransactions writes to sqldb, a SHYFT postgres instance
func swriteTransactions(db ethdb.SDatabase, tx *types.Transaction, blockHash common.Hash, blockNumber string, receipts []*types.Receipt, age time.Time, gasLimit uint64) error {
	var isContract bool
	var statusFromReciept, toAddr string
	var contractAddressFromReciept common.Address
	if tx.To() == nil {
		for _, receipt := range receipts {
			statusReciept := (*types.ReceiptForStorage)(receipt).Status
			contractAddressFromReciept = (*types.ReceiptForStorage)(receipt).ContractAddress
			switch {
			case statusReciept == 0:
				statusFromReciept = "FAIL"
			case statusReciept == 1:
				statusFromReciept = "SUCCESS"
			}
		}
		isContract = true
		tempAddr := &contractAddressFromReciept
		toAddr = tempAddr.String()
	} else {
		isContract = false
		for _, receipt := range receipts {
			statusReciept := (*types.ReceiptForStorage)(receipt).Status
			switch {
			case statusReciept == 0:
				statusFromReciept = "FAIL"
			case statusReciept == 1:
				statusFromReciept = "SUCCESS"
			}
		}
		toAddr = tx.To().String()
	}

	txData := stypes.ShyftTxEntryPretty{
		TxHash:      tx.Hash().Hex(),
		From:        tx.From().Hex(),
		To:          toAddr,
		BlockHash:   blockHash.Hex(),
		BlockNumber: blockNumber,
		Amount:      tx.Value().String(),
		Cost:        tx.Cost().String(),
		GasPrice:    tx.GasPrice().Uint64(),
		GasLimit:    gasLimit,
		Gas:         tx.Gas(),
		Nonce:       tx.Nonce(),
		Age:         age,
		Data:        tx.Data(),
		Status:      statusFromReciept,
		IsContract:  isContract,
	}
	isContractCheck := db.IsContract(txData.To)
	if isContractCheck {
		db.InsertTx(txData)
		//Runs necessary functions for tracing internal transactions through tracers.go
		if flag.Lookup("test.v") == nil {
			_, err := InternalTracker.TraceTransaction(tx.Hash(), blockHash)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		//Inserts Tx into DB
		db.InsertTx(txData)
	}
	return nil
}

// @NOTE: This function is extremely complex and requires heavy testing and knowdlege of edge cases:
// uncle blocks, account balance updates based on reorgs, diverges that get dropped.
// Reason for this is because the accounts are not deterministic like the block and tx hashes.
// @TODO: Calculate reorg - Determine how we include in acctblocks
func swriteMinerRewards(db ethdb.SDatabase, block *types.Block) string {
	minerAddr := block.Coinbase().String()
	shyftConduitAddress := common.HexToAddress("9db76b4bbaea76dfda4552b7b9d4e9d43abc55fd").String()
	// Calculate the total gas used in the block
	totalGas := new(big.Int)
	for _, tx := range block.Transactions() {
		totalGas.Add(totalGas, new(big.Int).Mul(tx.GasPrice(), new(big.Int).SetUint64(tx.Gas())))
	}

	blockHash := block.Hash().Hex()
	totalMinerReward := totalGas.Add(totalGas, Rewards.ShyftMinerBlockReward_v1)

	// References:
	// https://ethereum.stackexchange.com/questions/27172/different-uncles-reward
	// line 551 in consensus.go (go-empyrean/consensus/ethash/consensus.go)
	// Some weird constants to avoid constant memory allocs for them.
	var big8 = big.NewInt(8)
	var uncleRewards []*big.Int
	var uncleAddrs []string

	// uncleReward is overwritten after each iteration
	// Based on calculation in consensus.go accumulateRewards()
	uncleReward := new(big.Int)
	for _, uncle := range block.Uncles() {
		uncleReward.Add(uncle.Number, big8)
		uncleReward.Sub(uncleReward, block.Number())
		uncleReward.Mul(uncleReward, Rewards.ShyftMinerBlockReward_v1)
		uncleReward.Div(uncleReward, big8)
		uncleRewards = append(uncleRewards, uncleReward)
		uncleAddrs = append(uncleAddrs, uncle.Coinbase.String())
	}
	db.UpdateMinerAccount(minerAddr, blockHash, totalMinerReward)
	db.UpdateMinerAccount(shyftConduitAddress, blockHash, Rewards.ShyftNetworkBlockReward_v1)
	var uncRewards = new(big.Int)
	for i := 0; i < len(uncleAddrs); i++ {
		_ = uncleRewards[i]
		db.UpdateMinerAccount(uncleAddrs[i], blockHash, uncleRewards[i])
	}

	fullRewardValue := new(big.Int)
	fullRewardValue.Add(totalMinerReward, Rewards.ShyftNetworkBlockReward_v1)
	fullRewardValue.Add(fullRewardValue, uncRewards)

	return fullRewardValue.String()
}
