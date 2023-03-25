package whalealertapi

import (
	"fmt"
)

// Documentation can be found here:
// https://docs.whale-alert.io/

// Blockchain structure keeps info about blockchain name, available symbols, and blockchain status
type Blockchain struct {
	Name    string   `json:"name"`
	Symbols []string `json:"symbols"`
	Status  string   `json:"status"`
}

// Owner keeps data about wallet address owner
type Owner struct {
	Address   string `json:"address"`
	Owner     string `json:"owner"`
	OwnerType string `json:"owner_type"`
}

// Transaction keeps data about transaction which occured
type Transaction struct {
	Blockchain       string  `json:"blockchain"`
	Symbol           string  `json:"symbol"`
	TransactionType  string  `json:"transaction_type"`
	Hash             string  `json:"hash"`
	From             Owner   `json:"from"`
	To               Owner   `json:"to"`
	Timestamp        uint    `json:"timestamp"`
	Amount           float64 `json:"amount"`
	AmountUSD        float64 `json:"amount_usd"`
	TransactionCount uint    `json:"transaction_count"`
	ID               string  `json:"id"`
}

// TransactionResponse is returned when /transactions endpoint returns 200
type TransactionsResponse struct {
	Result       string        `json:"result"`
	Cursor       string        `json:"cursor"`
	Count        uint          `json:"count"`
	Transactions []Transaction `json:"transactions"`
}

// TransactionResponse is returned when /transaction endpoint returns 200
type TransactionResponse struct {
	Result       string        `json:"result"`
	Count        uint          `json:"count"`
	Transactions []Transaction `json:"transactions"`
}

// StatusResponse is returned when /status endpoint returns 200
type StatusResponse struct {
	Result          string       `json:"result"`
	BlockchainCount uint         `json:"blockchain_count"`
	Blockchains     []Blockchain `json:"blockchains"`
}

// ErrorResponse is returned when API reports an error
type ErrorResponse struct {
	Message string `json:"message"`
	Result  string `json:"result"`
	Err     error  `json:"error,omitempty"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf(e.Message)
}

type TransactionsRequest struct {
	Start    uint   `arg:"start"`
	End      uint   `arg:"end"`
	Cursor   string `arg:"cursor"`
	MinValue uint   `arg:"min_value"`
	Limit    uint   `arg:"limit"`
	Currency string `arg:"currency"`
}
