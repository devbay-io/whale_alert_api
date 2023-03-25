# WhaleAlert API Client

[![Go Coverage](https://github.com/devbay-io/whale_alert_api/wiki/coverage.svg)](https://raw.githack.com/wiki/devbay-io/whale_alert_api/coverage.html)

This is a Go API client for the WhaleAlert API.

## Usage

To use this API client, create a new instance of `WhaleAlertAPI` with the `New()` function, and then configure it using the various `With*` methods. Once you have configured the API client, you can make API requests using the `Status()`, `Transaction()` and `Transactions()` methods.

## Example

Here's an example of how to use this API client:

```golang
api := New().WithDefaultURL().WithAccessKey("your_api_key")

// Get the status of the API
status, err := api.Status()
if err != nil {
    log.Fatalf("Error getting API status: %v", err)
}
log.Printf("API status: %v", status)

// Get details about a transaction
transaction, err := api.Transaction("bitcoin", "tx_hash")
if err != nil {
    log.Fatalf("Error getting transaction details: %v", err)
}
log.Printf("Transaction details: %v", transaction)

// Get a list of transactions starting from a specific timestamp
start := uint(1631141147)
transactions, err := api.Transactions(start, TransactionsRequest{})
if err != nil {
    log.Fatalf("Error getting transactions: %v", err)
}
log.Printf("Transactions: %v", transactions)

```

## API Methods

### New()

`func New() *WhaleAlertAPI`

Creates a new instance of the API client.

### WithDefaultURL()

`func (api *WhaleAlertAPI) WithDefaultURL() *WhaleAlertAPI`

Sets the API URL to the default value of "https://api.whale-alert.io/v1". Returns the modified API client instance.

### WithCustomURL(url string)

`func (api *WhaleAlertAPI) WithCustomURL(url string) *WhaleAlertAPI`

Sets the API URL to a custom value. Returns the modified API client instance.

### WithAccessKey(key string)

`func (api *WhaleAlertAPI) WithAccessKey(key string) *WhaleAlertAPI`

Sets the API access key to a custom value. Returns the modified API client instance.

### Status()

`func (api WhaleAlertAPI) Status() (*StatusResponse, error)`

Retrieves the status of the WhaleAlert API. Returns a `StatusResponse` object and an error if the request fails.

### Transaction(blockchain, hash string)

`func (api WhaleAlertAPI) Transaction(blockchain, hash string)(*TransactionResponse, error)`

Retrieves details about a transaction with the specified blockchain and hash. Returns a `TransactionResponse` object and an error if the request fails.

### Transactions(start uint, args TransactionsRequest)

`func (api WhaleAlertAPI) Transactions(start uint, args TransactionsRequest) (*TransactionsResponse, error)`

Retrieves a list of transactions, starting from the specified Unix timestamp start, and using the provided `TransactionsRequest` arguments. Returns a TransactionsResponse object and an error if the request fails.

## License

This project is licensed under the MIT License - see the [LICENSE](/LICENSE) file for details.

## Authors

1. Mikołaj Mikołajczyk - code author
2. This README was generated by ChatGPT based on content of `api.go` file.
