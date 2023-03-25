package whalealertapi

import (
	"fmt"
	"net/http"
)

type WhaleAlertAPI struct {
	url    string
	key    string
	client *http.Client
}

func New() *WhaleAlertAPI {
	return &WhaleAlertAPI{
		client: &http.Client{},
	}
}

func (api *WhaleAlertAPI) WithDefaultURL() *WhaleAlertAPI {
	api.url = "https://api.whale-alert.io/v1"
	return api
}

func (api *WhaleAlertAPI) WithCustomURL(url string) *WhaleAlertAPI {
	api.url = url
	return api
}

func (api *WhaleAlertAPI) WithAccessKey(key string) *WhaleAlertAPI {
	api.key = key
	return api
}

func (api WhaleAlertAPI) Status() (*StatusResponse, error) {
	res, err := get[StatusResponse](api.client, api.url, api.key, "/status", []APIArgument{})
	return res, err
}

func (api WhaleAlertAPI) Transaction(blockchain, hash string) (*TransactionResponse, error) {
	if blockchain == "" || hash == "" {
		return nil, fmt.Errorf("blockchain and hash are required")
	}
	endpoint := fmt.Sprintf("/transaction/%s/%s", blockchain, hash)
	res, err := get[TransactionResponse](api.client, api.url, api.key, endpoint, []APIArgument{})
	return res, err
}

func (api WhaleAlertAPI) Transactions(start uint, args TransactionsRequest) (*TransactionsResponse, error) {
	if start <= 0 {
		return nil, fmt.Errorf("start must be greater than 0")
	}
	args.Start = start
	res, err := get[TransactionsResponse](api.client, api.url, api.key, "/transactions", args.toAPIArguments())
	return res, err
}
