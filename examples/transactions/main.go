package main

import (
	"fmt"
	"os"
	"strconv"

	whalealertapi "github.com/devbay-io/whale_alert_api"
)

// Example transaction output
// You need to specify following environment variables to run this example:
// WHALE_ALERT_API_KEY - your api key
// START_TIMESTAMP - unix timestamp

var (
	AccessKey      string
	StartTimestamp uint64
)

func init() {
	AccessKey = os.Getenv("WHALE_ALERT_API_KEY")
	var err error
	StartTimestamp, err = strconv.ParseUint(os.Getenv("START_TIMESTAMP"), 10, 64)
	if err != nil {
		panic(err)
	}
}

func main() {

	api := whalealertapi.New().WithDefaultURL().WithAccessKey(AccessKey)
	transactions, err := api.Transactions(uint(StartTimestamp), whalealertapi.TransactionsRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("This example shows how to request transactions data. You need to specify WHALE_ALERT_API_KEY, \n and START_TIMESTAMP environment variables to run this example.\n")

	fmt.Printf("Result: %v\n", transactions.Result)
	fmt.Printf("Count of transactions: %v and api cursor: %v\n", transactions.Count, transactions.Cursor)
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
