package shyftdb

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/syndtr/goleveldb/leveldb"

	"database/sql"

	_ "github.com/lib/pq"
)

//SBlock type
type SBlock struct {
	hash     string
	coinbase string
	number   string
}

//blockRes struct
type blockRes struct {
	hash     string
	coinbase string
	number   string
	Blocks   []SBlock
}

type txRes struct {
	txHash    string
	to_addr   string
	from_addr string
	blockHash string
	amount    string
	gasPrice  string
	gas       uint64
	nonce     uint64
	data      []byte
	TxEntry   []ShyftTxEntryPretty
}

//ShyftTxEntry structure
type ShyftTxEntry struct {
	TxHash    common.Hash
	To        *common.Address
	From      *common.Address
	BlockHash string
	Amount    *big.Int
	GasPrice  *big.Int
	Gas       uint64
	Nonce     uint64
	Data      []byte
}

type ShyftTxEntryPretty struct {
	TxHash    string
	To        string
	From      string
	BlockHash string
	Amount    string
	GasPrice  string
	Gas       uint64
	Nonce     uint64
	Data      []byte
}

type ShyftAccountEntry struct {
	Balance *big.Int
	Txs     []string
}

//WriteBlock writes to block info to sql db
func WriteBlock(sqldb *sql.DB, block *types.Block) error {
	coinbase := block.Header().Coinbase.String()
	number := block.Header().Number.String()

	sqlStatement := `INSERT INTO blocks(hash, coinbase, number) VALUES(($1), ($2), ($3)) RETURNING number`
	qerr := sqldb.QueryRow(sqlStatement, block.Header().Hash().Hex(), coinbase, number).Scan(&number) //.Scan(&fun)
	if qerr != nil {
		panic(qerr)
	}

	if block.Transactions().Len() > 0 {
		for _, tx := range block.Transactions() {
			WriteTransactions(sqldb, tx, block.Header().Hash())
			//tx_bytes[i] = tx.Hash().Bytes()
		}
	}
	return nil
}

//WriteTransactions writes to sqldb
func WriteTransactions(sqldb *sql.DB, tx *types.Transaction, blockHash common.Hash) error {
	txData := ShyftTxEntry{
		TxHash:    tx.Hash(),
		From:      tx.From(),
		To:        tx.To(),
		BlockHash: blockHash.Hex(),
		Amount:    tx.Value(),
		GasPrice:  tx.GasPrice(),
		Gas:       tx.Gas(),
		Nonce:     tx.Nonce(),
		Data:      tx.Data(),
	}

	txHash := txData.TxHash.Hex()
	from := txData.From.Hex()
	to := txData.To.Hex()
	blockHasher := txData.BlockHash
	amount := txData.Amount.String()
	gasPrice := txData.GasPrice.String()
	nonce := txData.Nonce
	gas := txData.Gas
	data := txData.Data

	sqlStatement := `INSERT INTO txs(txhash, to_addr, from_addr, blockhash, amount, gasprice, gas, nonce, data) VALUES(($1), ($2), ($3), ($4), ($5), ($6), ($7), ($8), ($9)) RETURNING nonce`
	qerr := sqldb.QueryRow(sqlStatement, txHash, to, from, blockHasher, amount, gasPrice, gas, nonce, data).Scan(&nonce)
	if qerr != nil {
		panic(qerr)
	}
	return nil
}

//WriteAccountBalances(db, tx)

// func WriteFromBalance(db *leveldb.DB, tx *types.Transaction) {
// 	key := append([]byte("acc-")[:], tx.From().Hash().Bytes()[:]...)
// 	// The from (sender) addr must have balance. If it fails to retrieve there is a bigger issue.
// 	retrievedData, err := db.Get(key, nil)
// 	if err != nil {
// 		log.Crit("From MUST have eth and no record found", "err", err)
// 	}
// 	var decodedData ShyftAccountEntry
// 	d := gob.NewDecoder(bytes.NewBuffer(retrievedData))
// 	if err := d.Decode(&decodedData); err != nil {
// 		log.Crit("Failed to decode From data:", "err", err)
// 	}
// 	decodedData.Balance.Sub(decodedData.Balance, tx.Value())
// 	decodedData.Txs = append(decodedData.Txs, tx.Hash().String())
// 	// Encode updated data
// 	var encodedData bytes.Buffer
// 	encoder := gob.NewEncoder(&encodedData)
// 	if err := encoder.Encode(decodedData); err != nil {
// 		log.Crit("Faild to encode From Account data", "err", err)
// 	}
// 	if err := db.Put(key, encodedData.Bytes(), nil); err != nil {
// 		log.Crit("Could not write the From account data", "err", err)
// 	}
// }

func WriteToBalance(db *leveldb.DB, tx *types.Transaction) {
	key := append([]byte("acc-")[:], tx.To().Hash().Bytes()[:]...)
	var txs []string

	retrievedData, err := db.Get(key, nil)
	if err != nil {
		accData := ShyftAccountEntry{
			Balance: tx.Value(),
			Txs:     append(txs, tx.Hash().String()),
		}
		var encodedData bytes.Buffer
		encoder := gob.NewEncoder(&encodedData)
		if err := encoder.Encode(accData); err != nil {
			log.Crit("Faild to encode To Account data", "err", err)
		}
		if err := db.Put(key, encodedData.Bytes(), nil); err != nil {
			log.Crit("Could not write the TO account's first tx", "err", err)
		}
	}
	var decodedData ShyftAccountEntry
	d := gob.NewDecoder(bytes.NewBuffer(retrievedData))
	if err := d.Decode(&decodedData); err != nil {
		log.Crit("Failed to decode To account data:", "err", err)
	}
	decodedData.Balance.Add(decodedData.Balance, tx.Value())
	decodedData.Txs = append(decodedData.Txs, tx.Hash().String())
	// Encode updated data
	var encodedData bytes.Buffer
	encoder := gob.NewEncoder(&encodedData)
	if err := encoder.Encode(decodedData); err != nil {
		log.Crit("Faild to encode To Account data", "err", err)
	}
	if err := db.Put(key, encodedData.Bytes(), nil); err != nil {
		log.Crit("Could not write the To account data", "err", err)
	}
}

// @NOTE: This function is extremely complex and requires heavy testing and knowdlege of edge cases:
// uncle blocks, account balance updates based on reorgs, diverges that get dropped.
// Reason for this is because the accounts are not deterministic like the block and tx hashes.
// @TODO: Calculate reward if there are uncles
// @TODO: Calculate mining reward (most likely retrieve higher up in the operations)
// @TODO: Calculate reorg
func WriteMinerReward(db *leveldb.DB, block *types.Block) {
	var totalGas *big.Int
	var txs []string
	key := append([]byte("acc-")[:], block.Coinbase().Hash().Bytes()[:]...)
	for _, tx := range block.Transactions() {
		totalGas.Add(totalGas, new(big.Int).Mul(tx.GasPrice(), new(big.Int).SetUint64(tx.Gas())))
	}
	retrievedData, err := db.Get(key, nil)
	if err != nil {
		// Assume time this account has had a tx
		// Balacne is exclusively minerreward + total gas from the block b/c no prior evm activity
		// Txs would be empty because they have not had any transactions on the EVM
		// @TODO: Calc mining reward
		//balance := totalGas.Add(totalGas, MINING_REWARD)
		balance := totalGas
		accData := ShyftAccountEntry{
			Balance: balance,
			Txs:     txs,
		}
		var encodedData bytes.Buffer
		encoder := gob.NewEncoder(&encodedData)
		if err := encoder.Encode(accData); err != nil {
			log.Crit("Faild to encode Miner Account data", "err", err)
		}
		if err := db.Put(key, encodedData.Bytes(), nil); err != nil {
			log.Crit("Could not write the miner's first tx", "err", err)
		}
	} else {
		// The account has already have previous data stored due to activity in the EVM
		// Decode the data to update balance
		var decodedData ShyftAccountEntry
		d := gob.NewDecoder(bytes.NewBuffer(retrievedData))
		if err := d.Decode(&decodedData); err != nil {
			log.Crit("Failed to decode miner data:", "err", err)
		}
		// Write new balance
		// @TODO: Calc mining reward
		// decodedData.Balance.Add(decodedData.Balance, totalGas.Add(totalGas, MINING_REWARD)))
		decodedData.Balance.Add(decodedData.Balance, totalGas)
		// Encode the data to be written back to the db
		var encodedData bytes.Buffer
		encoder := gob.NewEncoder(&encodedData)
		if err := encoder.Encode(decodedData); err != nil {
			log.Crit("Faild to encode Miner Account data", "err", err)
		}
		// Write newly encoded data back to the db
		if err := db.Put(key, encodedData.Bytes(), nil); err != nil {
			log.Crit("Could not update miner account data", "err", err)
		}
	}
}

///////////
// Getters
//////////
//GetAllBlocks returns []SBlock blocks for API
func GetAllBlocks(sqldb *sql.DB) []SBlock {
	var arr blockRes
	rows, err := sqldb.Query(`
		SELECT
			number,
			hash,
			coinbase
		FROM blocks`)
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	for rows.Next() {
		var num string
		var hash string
		var coinbase string
		err = rows.Scan(
			&num,
			&hash,
			&coinbase,
		)
		arr.Blocks = append(arr.Blocks, SBlock{
			hash:     hash,
			number:   num,
			coinbase: coinbase,
		})
	}
	return arr.Blocks
}

//GetBlock queries to send single block info
//TODO provide blockHash arg passed from handler.go
func GetBlock(sqldb *sql.DB) []SBlock {
	var arr blockRes
	sqlStatement := `SELECT * FROM blocks WHERE number=$1;`
	row := sqldb.QueryRow(sqlStatement, 3)
	var num string
	var hash string
	var coinbase string
	err := row.Scan(&num, &hash, &coinbase)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(num, hash, coinbase)
	default:
		panic(err)
	}
	arr.Blocks = append(arr.Blocks, SBlock{
		hash:     hash,
		number:   num,
		coinbase: coinbase,
	})
	return arr.Blocks
}

//GetAllTransactions getter fn for API
func GetAllTransactions(sqldb *sql.DB) []ShyftTxEntryPretty {
	var arr txRes
	rows, err := sqldb.Query(`
		SELECT
			txhash,
			to_addr,
			from_addr,
			blockhash,
			amount,
			gasprice,
			gas,
			nonce
		FROM txs`)
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	for rows.Next() {
		var txhash string
		var to_addr string
		var from_addr string
		var blockhash string
		var amount string
		var gasprice string
		var gas uint64
		var nonce uint64
		err = rows.Scan(
			&txhash,
			&to_addr,
			&from_addr,
			&blockhash,
			&amount,
			&gasprice,
			&gas,
			&nonce,
		)
		arr.TxEntry = append(arr.TxEntry, ShyftTxEntryPretty{
			TxHash:    txhash,
			To:        to_addr,
			From:      from_addr,
			BlockHash: blockhash,
			Amount:    amount,
			GasPrice:  gasprice,
			Gas:       gas,
			Nonce:     nonce,
		})
	}
	return arr.TxEntry
}

//GetTransaction fn returns single tx
func GetTransaction(sqldb *sql.DB) []ShyftTxEntryPretty {
	var arr txRes
	sqlStatement := `SELECT * FROM txs WHERE nonce=$1;`
	row := sqldb.QueryRow(sqlStatement, 7)
	var txhash string
	var to_addr string
	var from_addr string
	var blockhash string
	var amount string
	var gasprice string
	var gas uint64
	var nonce uint64
	_ = row.Scan(&txhash, &to_addr, &from_addr, &blockhash, &amount, &gasprice, &gas, &nonce)

	// switch err {
	// case sql.ErrNoRows:
	// 	fmt.Println("No rows were returned!")
	// case nil:
	// 	fmt.Println(txhash, to_addr, from_addr, blockhash, amount, gasprice, gas, nonce)
	// default:
	// 	panic(err)
	// }

	arr.TxEntry = append(arr.TxEntry, ShyftTxEntryPretty{
		TxHash:    txhash,
		To:        to_addr,
		From:      from_addr,
		BlockHash: blockhash,
		Amount:    amount,
		GasPrice:  gasprice,
		Gas:       gas,
		Nonce:     nonce,
	})

	return arr.TxEntry
}