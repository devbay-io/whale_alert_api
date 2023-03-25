package main

import (
	"fmt"
	"os"
	"strings"

	whalealertapi "github.com/devbay-io/whale_alert_api"
)

var AccessKey string

func init() {
	AccessKey = os.Getenv("WHALE_ALERT_API_KEY")
}

func main() {

	api := whalealertapi.New().WithDefaultURL().WithAccessKey(AccessKey)
	status, err := api.Status()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status example shows current API status below.\n")
	fmt.Printf("Request result: %s\n", status.Result)
	fmt.Printf("Blockchain count: %d\n", status.BlockchainCount)
	fmt.Printf("Blockchain data:\n")
	for _, blockchain := range status.Blockchains {
		fmt.Printf("Blockchain: %s, status: %s, symbols: %s\n", blockchain.Name, blockchain.Status, strings.Join(blockchain.Symbols, ","))
	}
}
