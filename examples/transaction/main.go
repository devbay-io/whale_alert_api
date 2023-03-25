package main

import (
	"fmt"
	"os"

	whalealertapi "github.com/devbay-io/whale_alert_api"
)

// Example transaction output
// You need to specify following environment variables to run this example:
// WHALE_ALERT_API_KEY - your api key
// TRANSACTION_HASH - correct transaction hash
// BLOCKCHAIN - correct blockchain
// example transaction hash b13a7ba1d0232779fa8465715a5401a7b145271a1146415d34f34ee2dc86ad48 and blockchain ethereum

var (
	AccessKey       string
	TransactionHash string
	Blockchain      string
)

func init() {
	AccessKey = os.Getenv("WHALE_ALERT_API_KEY")
	TransactionHash = os.Getenv("TRANSACTION_HASH")
	Blockchain = os.Getenv("BLOCKCHAIN")
}

func main() {

	api := whalealertapi.New().WithDefaultURL().WithAccessKey(AccessKey)
	transactions, err := api.Transaction(Blockchain, TransactionHash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("This example shows how to request transaction data. You need to specify WHALE_ALERT_API_KEY, \nTRANSACTION_HASH and BLOCKCHAIN environment variables to run this example.\n")

	fmt.Printf("Result: %v\n", transactions.Result)
	fmt.Printf("Count of transactions: %v\n", transactions.Count)
	for i, transaction := range transactions.Transactions {
		fmt.Printf("Transaction number: %d\n", i+1)
		fmt.Printf("Transaction hash %v\n", transaction.Hash)
		fmt.Printf("Transaction blockchain: %v and symbol: %v\n", transaction.Blockchain, transaction.Symbol)
		fmt.Printf("Transaction amount: %v and in USD: %v\n", transaction.Amount, transaction.AmountUSD)
		fmt.Printf("Transaction date: %v\n", transaction.Timestamp)
		fmt.Printf("Transaction from: %v\n", transaction.From.Address)
		fmt.Printf("Transaction to: %v\n", transaction.To.Address)
	}
}
