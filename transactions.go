package walletapi

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Transfer - represents a transfer object
type Transfer struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

// TransactionResult - represents the result of a sent / prepared transaction
type TransactionResult struct {
	TransactionHash  string `json:"transactionHash"`
	Fee              uint64 `json:"fee"`
	RelayedToNetwork bool   `json:"relayedToNetwork"`
}

// Transaction - represents a transaction object
type Transaction struct {
	BlockHeight           uint64     `json:"blockHeight"`
	Fee                   uint64     `json:"fee"`
	FeePerByte            uint64     `json:"feePerByte"`
	Hash                  string     `json:"hash"`
	IsCoinbaseTransaction bool       `json:"isCoinbaseTransaction"`
	PaymentID             string     `json:"paymentID"`
	Timestamp             uint64     `json:"timestamp"`
	UnlockTime            uint64     `json:"unlockTime"`
	Transfers             []Transfer `json:"transfers"`
}

// Transactions - represents a transactions object
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
	Transaction  Transaction   `json:"transaction"`
}

// GetAllTransactions - gets all the transactions in the wallet container
func (wAPI WalletAPI) GetAllTransactions() (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions",
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetTransactionByHash - gets the transaction with the given hash in the wallet container
func (wAPI WalletAPI) GetTransactionByHash(hash string) (tx *Transaction, err error) {
	var txs Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[404])
		}
	}()

	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/hash/"+hash,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &txs)
		if err != nil {
			panic(err)
		}
	}

	return &txs.Transaction, err
}

// GetUnconfirmedTransactions - gets all unconfirmed outgoing transactions
func (wAPI WalletAPI) GetUnconfirmedTransactions() (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/unconfirmed",
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetUnconfirmedTransactionsByAddress - gets all unconfirmed outgoing transactions for a given address
func (wAPI WalletAPI) GetUnconfirmedTransactionsByAddress(address string) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/unconfirmed/"+address,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetTransactionsByStartHeight - gets 1000 transactions for the wallet starting at startHeight
func (wAPI WalletAPI) GetTransactionsByStartHeight(startHeight uint64) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	start := strconv.FormatUint(startHeight, 10)
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/"+start,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetTransactionsInRange - gets transactions for the wallet given a range of block heights
func (wAPI WalletAPI) GetTransactionsInRange(start, end uint64) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	low := strconv.FormatUint(start, 10)
	high := strconv.FormatUint(end, 10)
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/"+low+"/"+high,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetAddressTransactionsByStartHeight - gets 1000 transactions for the address starting at startHeight
func (wAPI WalletAPI) GetAddressTransactionsByStartHeight(address string, startHeight uint64) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	start := strconv.FormatUint(startHeight, 10)
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/address/"+address+"/"+start,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// GetAddressTransactionsInRange - gets transactions for the address given a range of block heights
func (wAPI WalletAPI) GetAddressTransactionsInRange(address string, start, end uint64) (txs *[]Transaction, err error) {
	var tx Transactions
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[500])
		}
	}()

	low := strconv.FormatUint(start, 10)
	high := strconv.FormatUint(end, 10)
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/address/"+address+"/"+low+"/"+high,
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &tx)
		if err != nil {
			panic(err)
		}
	}

	return &tx.Transactions, err
}

// SendTransactionBasic - sends a transaction
func (wAPI WalletAPI) SendTransactionBasic(destination, paymentID string, amount uint64) (string, error) {
	var txHash string

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/send/basic",
		makeJSONString(map[string]interface{}{
			"destination": destination,
			"amount":      amount,
			"paymentID":   paymentID,
		}),
	)

	if err == nil {
		txHash = (*resp)["transactionHash"].(string)
	}

	return txHash, err
}

// SendTransactionAdvanced - sends a transaction
func (wAPI WalletAPI) SendTransactionAdvanced(
	destinations []map[string]interface{},
	mixin, fee, feePerByte, sourceAddresses, paymentID, changeAddress, unlockTime interface{}) (tx *TransactionResult, err error) {
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[400])
		}
	}()

	body := map[string]interface{}{
		"destinations": destinations,
	}

	if mixin != nil {
		body["mixin"] = mixin.(uint64)
	}
	if fee != nil {
		body["fee"] = fee.(uint64)
	}
	if feePerByte != nil {
		body["feePerByte"] = feePerByte.(uint64)
	}
	if sourceAddresses != nil {
		body["sourceAddresses"] = sourceAddresses.([]string)
	}
	if paymentID != nil {
		body["paymentID"] = paymentID.(string)
	}
	if unlockTime != nil {
		body["unlockTime"] = unlockTime.(uint64)
	}
	if changeAddress != nil {
		body["changeAddress"] = changeAddress.(string)
	}

	_, raw, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/send/advanced",
		makeJSONString(body),
	)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(*raw, &tx)

	return tx, err
}

// SendFusionBasic - sends a fusion transaction
func (wAPI WalletAPI) SendFusionBasic() (string, error) {
	var txHash string

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/send/fusion/basic",
		"",
	)

	if err == nil {
		txHash = (*resp)["transactionHash"].(string)
	}

	return txHash, err
}

// SendFusionAdvanced - sends a fusion transaction
func (wAPI WalletAPI) SendFusionAdvanced(sourceAddresses []string, destination string) (string, error) {
	var txHash string

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/send/fusion/advanced",
		makeJSONString(map[string]interface{}{
			"mixin":           MIXIN,
			"sourceAddresses": sourceAddresses,
			"destination":     destination,
		}),
	)

	if err == nil {
		txHash = (*resp)["transactionHash"].(string)
	}

	return txHash, err
}

// PrepareTransactionBasic - creates a transaction but does not relay it to the network
func (wAPI WalletAPI) PrepareTransactionBasic(destination string, amount uint64, paymentID string) (*TransactionResult, error) {
	var transaction TransactionResult

	_, raw, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/prepare/basic",
		makeJSONString(map[string]interface{}{
			"destination": destination,
			"amount":      amount,
			"paymentID":   paymentID,
		}),
	)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*raw, &transaction)

	return &transaction, err
}

// PrepareTransactionAdvanced - creates a transaction but does not relay it to the network
func (wAPI WalletAPI) PrepareTransactionAdvanced(
	destinations []map[string]interface{},
	mixin, fee, feePerByte, sourceAddresses, paymentID, changeAddress, unlockTime interface{}) (tx *TransactionResult, err error) {
	defer func() {
		if ok := recover(); ok != nil {
			err = errors.New(ERRORS[400])
		}
	}()

	body := map[string]interface{}{
		"destinations": destinations,
	}

	if mixin != nil {
		body["mixin"] = mixin.(uint64)
	}
	if fee != nil {
		body["fee"] = fee.(uint64)
	}
	if feePerByte != nil {
		body["feePerByte"] = feePerByte.(uint64)
	}
	if sourceAddresses != nil {
		body["sourceAddresses"] = sourceAddresses.([]string)
	}
	if paymentID != nil {
		body["paymentID"] = paymentID.(string)
	}
	if unlockTime != nil {
		body["unlockTime"] = unlockTime.(uint64)
	}
	if changeAddress != nil {
		body["changeAddress"] = changeAddress.(string)
	}

	_, raw, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/prepare/advanced",
		makeJSONString(body),
	)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(*raw, &tx)

	return tx, err
}

// SendPreparedTransaction - sends a previously created prepared transaction
func (wAPI WalletAPI) SendPreparedTransaction(hash string) (string, error) {
	var txHash string

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/transactions/prepared",
		makeJSONString(map[string]interface{}{
			"transactionHash": hash,
		}),
	)

	if err == nil {
		txHash = (*resp)["transactionHash"].(string)
	}

	return txHash, err
}

// DeletePreparedTransaction - removes a previously created prepared transaction
func (wAPI WalletAPI) DeletePreparedTransaction(hash string) error {
	_, _, err := wAPI.sendRequest(
		"DELETE",
		wAPI.Host+":"+wAPI.Port+"/transactions/prepared/"+hash,
		"",
	)

	return err
}

// GetTransactionPrivateKey - gets the private key of a transaction with the given hash
func (wAPI WalletAPI) GetTransactionPrivateKey(hash string) (string, error) {
	var privKey string

	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/transactions/privatekey/"+hash,
		"",
	)
	if err == nil {
		privKey = (*resp)["transactionPrivateKey"].(string)
	}

	return privKey, err
}
