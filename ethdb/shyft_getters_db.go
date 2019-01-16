package ethdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ShyftNetwork/go-empyrean/core/sTypes"
)

func SGetAllBlocksWithoutLimit() (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println("err")
		return "", err
	}
	var arr stypes.BlockRes
	var blockArr string

	rows, err := db.db.Queryx(`SELECT * FROM blocks ORDER BY number ASC`)
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var hash, coinbase, parentHash, uncleHash, difficulty, size, rewards, num string
		var gasUsed, gasLimit, nonce uint64
		var txCount, uncleCount int
		var age time.Time

		err = rows.Scan(
			&hash, &coinbase, &gasUsed, &gasLimit, &txCount, &uncleCount, &age, &parentHash, &uncleHash, &difficulty, &size, &nonce, &rewards, &num)

		arr.Blocks = append(arr.Blocks, stypes.SBlock{
			Hash:       hash,
			Coinbase:   coinbase,
			GasUsed:    gasUsed,
			GasLimit:   gasLimit,
			TxCount:    txCount,
			UncleCount: uncleCount,
			Age:        age,
			ParentHash: parentHash,
			UncleHash:  uncleHash,
			Difficulty: difficulty,
			Size:       size,
			Nonce:      nonce,
			Rewards:    rewards,
			Number:     num,
		})

		blocks, _ := json.Marshal(arr.Blocks)
		blocksFmt := string(blocks)
		blockArr = blocksFmt
	}
	return blockArr, nil
}

// GetAllBlocksLength returns number of records in blocks table
// This tells the client-UI how many pages to show for pagination
func GetAllBlocksLength() (int, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println("err")
		return 0, err
	}
	sqlStatement := `SELECT COUNT(*) FROM blocks`
	tx, _ := db.db.Begin()
	row := db.db.QueryRow(sqlStatement)
	tx.Commit()

	var count int
	row.Scan(
		&count)

	return count, nil
}

func SGetAllBlocks(limit string, offset string) (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println("err")
		return "", err
	}
	var arr stypes.BlockRes
	var blockArr string
	sqlStatement := `SELECT * FROM blocks ORDER BY number ASC LIMIT $1 OFFSET $2`
	rows, err := db.db.Queryx(sqlStatement, limit, offset)
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var hash, coinbase, parentHash, uncleHash, difficulty, size, rewards, num string
		var gasUsed, gasLimit, nonce uint64
		var txCount, uncleCount int
		var age time.Time

		err = rows.Scan(
			&hash, &coinbase, &gasUsed, &gasLimit, &txCount, &uncleCount, &age, &parentHash, &uncleHash, &difficulty, &size, &nonce, &rewards, &num)

		arr.Blocks = append(arr.Blocks, stypes.SBlock{
			Hash:       hash,
			Coinbase:   coinbase,
			GasUsed:    gasUsed,
			GasLimit:   gasLimit,
			TxCount:    txCount,
			UncleCount: uncleCount,
			Age:        age,
			ParentHash: parentHash,
			UncleHash:  uncleHash,
			Difficulty: difficulty,
			Size:       size,
			Nonce:      nonce,
			Rewards:    rewards,
			Number:     num,
		})

		blocks, _ := json.Marshal(arr.Blocks)
		blocksFmt := string(blocks)
		blockArr = blocksFmt
	}
	return blockArr, nil
}

//GetBlock queries to send single block info
//TODO provide blockHash arg passed from handler.go
func SGetBlock(blockNumber string) (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	sqlStatement := `SELECT * FROM blocks WHERE number=$1;`
	tx, _ := db.db.Begin()
	row := db.db.QueryRow(sqlStatement, blockNumber)
	tx.Commit()
	var hash, coinbase, parentHash, uncleHash, difficulty, size, rewards, num string
	var gasUsed, gasLimit, nonce uint64
	var txCount, uncleCount int
	var age time.Time

	row.Scan(
		&hash, &coinbase, &gasUsed, &gasLimit, &txCount, &uncleCount, &age, &parentHash, &uncleHash, &difficulty, &size, &nonce, &rewards, &num)

	block := stypes.SBlock{
		Hash:       hash,
		Coinbase:   coinbase,
		GasUsed:    gasUsed,
		GasLimit:   gasLimit,
		TxCount:    txCount,
		UncleCount: uncleCount,
		Age:        age,
		ParentHash: parentHash,
		UncleHash:  uncleHash,
		Difficulty: difficulty,
		Size:       size,
		Nonce:      nonce,
		Rewards:    rewards,
		Number:     num,
	}
	json, _ := json.Marshal(block)
	return string(json), nil
}

func SGetRecentBlock() (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	sqlStatement := `SELECT * FROM blocks WHERE number=(SELECT MAX(number) FROM blocks);`
	tx, _ := db.db.Begin()
	row := db.db.QueryRow(sqlStatement)
	tx.Commit()
	var hash, coinbase, parentHash, uncleHash, difficulty, size, rewards, num string
	var gasUsed, gasLimit, nonce uint64
	var txCount, uncleCount int
	var age time.Time

	row.Scan(
		&hash, &coinbase, &gasUsed, &gasLimit, &txCount, &uncleCount, &age, &parentHash, &uncleHash, &difficulty, &size, &nonce, &rewards, &num)

	block := stypes.SBlock{
		Hash:       hash,
		Coinbase:   coinbase,
		GasUsed:    gasUsed,
		GasLimit:   gasLimit,
		TxCount:    txCount,
		UncleCount: uncleCount,
		Age:        age,
		ParentHash: parentHash,
		UncleHash:  uncleHash,
		Difficulty: difficulty,
		Size:       size,
		Nonce:      nonce,
		Rewards:    rewards,
		Number:     num,
	}
	json, _ := json.Marshal(block)
	return string(json), nil
}

//func SGetRecentBlockHash() string {
//	sqldb, _ := DBConnection()
//	sqlStatement := `SELECT hash FROM blocks WHERE number=(SELECT MAX(number) FROM blocks);`
//	tx, _ := sqldb.Begin()
//	row := sqldb.QueryRow(sqlStatement)
//	tx.Commit()
//	var hash string
//
//	row.Scan(
//		&hash)
//
//	blockhash := stypes.BlockHash{
//		Hash: hash,
//	}
//	json, _ := json.Marshal(blockhash)
//	return string(json)
//}

func SGetAllTransactionsFromBlock(blockNumber string) (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var arr stypes.TxRes
	var txx string
	sqlStatement := `SELECT * FROM txs WHERE blocknumber=$1`
	tx, _ := db.db.Begin()
	rows, err := db.db.Query(sqlStatement, blockNumber)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	for rows.Next() {
		var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
		var gasprice, gas, gasLimit, nonce uint64
		var isContract bool
		var age time.Time
		var data []byte

		err = rows.Scan(
			&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data,
		)

		arr.TxEntry = append(arr.TxEntry, stypes.ShyftTxEntryPretty{
			TxHash:      txhash,
			To:          to_addr,
			From:        from_addr,
			BlockHash:   blockhash,
			BlockNumber: blocknumber,
			Amount:      amount,
			GasPrice:    gasprice,
			Gas:         gas,
			GasLimit:    gasLimit,
			Cost:        txfee,
			Nonce:       nonce,
			Status:      status,
			IsContract:  isContract,
			Age:         age,
			Data:        data,
		})

		txData, _ := json.Marshal(arr.TxEntry)
		newtx := string(txData)
		txx = newtx
	}
	return txx, nil
}

func SGetAllBlocksMinedByAddress(coinbase string) (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	var arr stypes.BlockRes
	var blockArr string
	sqlStatement := `SELECT * FROM blocks WHERE coinbase=$1`
	tx, _ := db.db.Begin()
	rows, err := db.db.Query(sqlStatement, coinbase)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	for rows.Next() {
		var hash, coinbase, parentHash, uncleHash, difficulty, size, rewards, num string
		var gasUsed, gasLimit, nonce uint64
		var txCount, uncleCount int
		var age time.Time

		err = rows.Scan(
			&hash, &coinbase, &gasUsed, &gasLimit, &txCount, &uncleCount, &age, &parentHash, &uncleHash, &difficulty, &size, &nonce, &rewards, &num)

		arr.Blocks = append(arr.Blocks, stypes.SBlock{
			Hash:       hash,
			Coinbase:   coinbase,
			GasUsed:    gasUsed,
			GasLimit:   gasLimit,
			TxCount:    txCount,
			UncleCount: uncleCount,
			Age:        age,
			ParentHash: parentHash,
			UncleHash:  uncleHash,
			Difficulty: difficulty,
			Size:       size,
			Nonce:      nonce,
			Rewards:    rewards,
			Number:     num,
		})

		blocks, _ := json.Marshal(arr.Blocks)
		blocksFmt := string(blocks)
		blockArr = blocksFmt
	}
	return blockArr, nil
}

//GetAllTransactions getter fn for API
func SGetAllTransactions() (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var arr stypes.TxRes
	var txx string
	tx, _ := db.db.Begin()
	rows, err := db.db.Query(`SELECT * FROM txs`)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	for rows.Next() {
		var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
		var gasprice, gas, gasLimit, nonce uint64
		var isContract bool
		var age time.Time
		var data []byte

		err = rows.Scan(
			&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data,
		)

		arr.TxEntry = append(arr.TxEntry, stypes.ShyftTxEntryPretty{
			TxHash:      txhash,
			To:          to_addr,
			From:        from_addr,
			BlockHash:   blockhash,
			BlockNumber: blocknumber,
			Amount:      amount,
			GasPrice:    gasprice,
			Gas:         gas,
			GasLimit:    gasLimit,
			Cost:        txfee,
			Nonce:       nonce,
			Status:      status,
			IsContract:  isContract,
			Age:         age,
			Data:        data,
		})

		txData, _ := json.Marshal(arr.TxEntry)
		newtx := string(txData)
		txx = newtx
	}
	return txx, nil
}

//GetTransaction fn returns single tx
func SGetTransaction(txHash string) (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	sqlStatement := `SELECT * FROM txs WHERE txhash=$1;`
	tx, _ := db.db.Begin()
	row := db.db.QueryRow(sqlStatement, txHash)
	tx.Commit()
	var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
	var gasprice, gas, gasLimit, nonce uint64
	var isContract bool
	var age time.Time
	var data []byte

	row.Scan(
		&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data)

	txData := stypes.ShyftTxEntryPretty{
		TxHash:      txhash,
		To:          to_addr,
		From:        from_addr,
		BlockHash:   blockhash,
		BlockNumber: blocknumber,
		Amount:      amount,
		GasPrice:    gasprice,
		Gas:         gas,
		GasLimit:    gasLimit,
		Cost:        txfee,
		Nonce:       nonce,
		Status:      status,
		IsContract:  isContract,
		Age:         age,
		Data:        data,
	}
	json, _ := json.Marshal(txData)

	return string(json), nil
}

//GetAllAccounts returns all accounts and balances
func SGetAllAccounts() (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var array stypes.AccountRes
	var accountsArr, nonce string
	tx, _ := db.db.Begin()
	accs, err := db.db.Query(`
		SELECT
			addr,
			balance,
			nonce
		FROM accounts`)
	tx.Commit()
	if err != nil {
		fmt.Println(err)
	}

	defer accs.Close()

	for accs.Next() {
		var addr, balance string
		err = accs.Scan(
			&addr, &balance, &nonce,
		)

		array.AllAccounts = append(array.AllAccounts, stypes.SAccounts{
			Addr:         addr,
			Balance:      balance,
			AccountNonce: nonce,
		})

		accounts, _ := json.Marshal(array.AllAccounts)
		accountsFmt := string(accounts)
		accountsArr = accountsFmt
	}
	return accountsArr, nil
}

func InnerSGetAccount(db *SPGDatabase, address string) (stypes.SAccounts, bool) {
	sqlStatement := `SELECT * FROM accounts WHERE addr=$1;`
	var addr, balance, nonce string
	tx, _ := db.db.Begin()
	err := db.db.QueryRow(sqlStatement, address).Scan(&addr, &balance, &nonce)
	tx.Commit()
	if err == sql.ErrNoRows {
		return stypes.SAccounts{}, false
	} else {
		account := stypes.SAccounts{
			Addr:         addr,
			Balance:      balance,
			AccountNonce: nonce,
		}
		return account, true
	}
}

//GetAccount returns account balances
func SGetAccount(address string) (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var account, _ = InnerSGetAccount(db, address)
	json, _ := json.Marshal(account)
	return string(json), nil
}

//GetAccount returns account balances
func SGetAccountTxs(address string) (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var arr stypes.TxRes
	var txx string
	sqlStatement := `SELECT * FROM txs WHERE to_addr=$1 OR from_addr=$1;`
	tx, _ := db.db.Begin()
	rows, err := db.db.Query(sqlStatement, address)
	tx.Commit()
	if err != nil {
		fmt.Println("err", err)
	}
	defer rows.Close()
	for rows.Next() {
		var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
		var gasprice, gas, gasLimit, nonce uint64
		var isContract bool
		var age time.Time
		var data []byte

		err = rows.Scan(
			&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data,
		)

		arr.TxEntry = append(arr.TxEntry, stypes.ShyftTxEntryPretty{
			TxHash:      txhash,
			To:          to_addr,
			From:        from_addr,
			BlockHash:   blockhash,
			BlockNumber: blocknumber,
			Amount:      amount,
			GasPrice:    gasprice,
			Gas:         gas,
			GasLimit:    gasLimit,
			Cost:        txfee,
			Nonce:       nonce,
			Status:      status,
			IsContract:  isContract,
			Age:         age,
			Data:        data,
		})

		txData, _ := json.Marshal(arr.TxEntry)
		newtx := string(txData)
		txx = newtx
	}
	return txx, nil
}

//GetAllInternalTransactions getter fn for API
func SGetAllInternalTransactions() (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var arr stypes.InternalArray
	var internaltx string
	tx, _ := db.db.Begin()
	rows, err := db.db.Query(`SELECT * FROM internaltxs`)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	for rows.Next() {
		var txhash, blockhash, action, to_addr, from_addr, amount, input, output string
		var gas, gasUsed uint64
		var id int
		var age string

		err = rows.Scan(
			&id, &txhash, &blockhash, &action, &to_addr, &from_addr, &amount, &gas, &gasUsed, &age, &input, &output,
		)

		arr.InternalEntry = append(arr.InternalEntry, stypes.InteralWrite{
			ID:        id,
			Hash:      txhash,
			BlockHash: blockhash,
			Action:    action,
			To:        to_addr,
			From:      from_addr,
			Value:     amount,
			Gas:       gas,
			GasUsed:   gasUsed,
			Time:      age,
			Input:     input,
			Output:    output,
		})

		txData, _ := json.Marshal(arr.InternalEntry)
		newtx := string(txData)
		internaltx = newtx
	}
	return internaltx, nil
}

//GetInternalTransaction fn returns single tx
func SGetInternalTransaction(txHash string) (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	var arr stypes.InternalArray
	var internaltx string

	sqlStatement := `SELECT * FROM internaltxs WHERE txhash=$1;`
	tx, _ := db.db.Begin()
	rows, err := db.db.Query(sqlStatement, txHash)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	for rows.Next() {
		var txhash, blockhash, action, to_addr, from_addr, amount, input, output string
		var id int
		var gas, gasUsed uint64
		var age string

		err = rows.Scan(
			&id, &txhash, &blockhash, &action, &to_addr, &from_addr, &amount, &gas, &gasUsed, &age, &input, &output,
		)

		arr.InternalEntry = append(arr.InternalEntry, stypes.InteralWrite{
			ID:        id,
			Hash:      txhash,
			BlockHash: blockhash,
			Action:    action,
			To:        to_addr,
			From:      from_addr,
			Value:     amount,
			Gas:       gas,
			GasUsed:   gasUsed,
			Time:      age,
			Input:     input,
			Output:    output,
		})

		txData, _ := json.Marshal(arr.InternalEntry)
		newtx := string(txData)
		internaltx = newtx
	}
	return internaltx, nil
}

func SGetAllAccountBlocks() (string, error) {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var arr AccountBlockArray
	var accountBlockJSON string
	tx, _ := db.db.Begin()
	rows, err := db.db.Query(`SELECT * FROM accountblocks`)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	for rows.Next() {
		var acct, blockhash string
		var delta, txCount int64

		err = rows.Scan(
			&acct, &blockhash, &delta, &txCount,
		)

		arr.AccountBlocks = append(arr.AccountBlocks, AccountBlock{
			Acct:      acct,
			Blockhash: blockhash,
			Delta:     delta,
			TxCount:   txCount,
		})

		accountBlocks, _ := json.Marshal(arr.AccountBlocks)
		accountBlocksJSON := string(accountBlocks)
		accountBlockJSON = accountBlocksJSON
	}
	return accountBlockJSON, nil
}
