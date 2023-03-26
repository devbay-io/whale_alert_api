package whalealertapi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	whalealertapi "github.com/devbay-io/whale_alert_api_client"
)

func TestStatus(t *testing.T) {
	serverOK := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("X-WA-API-KEY") == "CORRECT" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result":"success","blockchain_count":15,"blockchains":[{"name":"ethereum","symbols":["ht","iht","paxg","san","okb","ankr","dusk","hex","husd","iotx","tusd","nexo","ocn","rsr","theta","qbit","bhpc","bzrx","kan","la","mda","usdt","chz","ppt","storj","uni","appc","gvt","hot","icn","nxm","bnt","drgn","hedg","mxm","shib","aoa","dcn","nas","cnd","gnx","iost","thx","brc","cro","dgtx","powr","tel","link","pnk","sand","aave","busd","dx","eth","inb","susd","loom","omg","vest","yfi","ren","ttc","agi","lky","qnt","sub","thr","abt","btc","poly","r","rhoc","sushi","aion","bix","btm","eurs","ftt","credo","ethos","wax","amp","kcs","pax","usdc","vibe","gnt","lrc","storm","wtc","xaut","bal","dgd","eurt","osa","xyo","cennz","elf","knc","pay","rep","cmt","salt","sxdt","cosm","kin","lend","nuls","tomo","qkc","srn","gusd","wbtc","yfii","cdai","cel","dai","eng","trac","zrx","brd","grt","man","npxs","snt","mkr","tct","bczero","ecoreal","evx","icx","mco","tnb","ampl","edr","mgo","sxp","meta","ae","bat","cnht","cvc","edo","veri","bnb","mft","qash","rlc","snx","req","skl","wic","dac","dent","fun","mana","mtl","zil","enj","leo","nec","xin","uma","ctxc","hpt","matic","ncash","trx"],"status":"connected"},{"name":"neo","symbols":["neo","gas"],"status":"connected"},{"name":"binancechain","symbols":["bnb"],"status":"connected"},{"name":"bitcoin","symbols":["usdt","btc","eurt"],"status":"connected"},{"name":"hive","symbols":["hive",""],"status":"connected"},{"name":"icon","symbols":["icx"],"status":"connected"},{"name":"cosmos","symbols":["atom"],"status":"connected"},{"name":"eos","symbols":["usdt","eos","leo"],"status":"connected"},{"name":"stellar","symbols":["mobi","","xlm","slt","repo"],"status":"connected"},{"name":"unknown","symbols":[""],"status":"connected"},{"name":"liquid","symbols":["usdt"],"status":"connected"},{"name":"tron","symbols":["usdt","btt","bnb","trx"],"status":"connected"},{"name":"steem","symbols":["steem"],"status":"connected"},{"name":"tezos","symbols":["xtz","amount"],"status":"connected"},{"name":"ripple","symbols":["","xrp"],"status":"connected"}]}`))
			return
		}
		if r.Header.Get("X-WA-API-KEY") == "MALFORMED_JSON" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"result":"success","block`))
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"result":"error","message":"invalid api_key"}`))
	}))

	api := whalealertapi.New().WithCustomURL(serverOK.URL).WithAccessKey("CORRECT")
	res, err := api.Status()
	if err != nil {
		t.Errorf("Expected OK got error: %s", err)
	}
	if res.BlockchainCount != 15 {
		t.Errorf("Expected %d got: %d", 15, res.BlockchainCount)
	}
	if res.Result != "success" {
		t.Errorf("Expected %s got: %s", "success", res.Result)
	}
	api.WithAccessKey("MALFORMED_JSON")
	res, err = api.Status()
	if err == nil {
		t.Errorf("Expected error")
	}
	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
	if err.Error() != "unexpected EOF" {
		t.Errorf("Expected %s got: %s", "unexpected EOF", err.Error())
	}
	api.WithAccessKey("UNAUTHORIZED")
	res, err = api.Status()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
	if err.Error() != "invalid api_key" {
		t.Errorf("Expected %s got: %s", "invalid api_key", err.Error())
	}
}

const (
	hashCorrect    = "b13a7ba1d0232779fa8465715a5401a7b145271a1146415d34f34ee2dc86ad48"
	hashIncorrect  = "779fa8465715a5401a7b145271a1146415d34f34ee2dc86ad48"
	ethChain       = "ethereum"
	btcChain       = "bitcoin"
	incorrectChain = "abc"
)

func TestTransaction(t *testing.T) {
	serverOK := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//When correct blockchain and hash
		if r.Header.Get("X-WA-API-KEY") == "CORRECT" && strings.Contains(r.URL.Path, ethChain) && strings.Contains(r.URL.Path, hashCorrect) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result":"success","count":2,"transactions":[{"blockchain":"ethereum","symbol":"usdt","id":"1990512547","transaction_type":"transfer","hash":"b13a7ba1d0232779fa8465715a5401a7b145271a1146415d34f34ee2dc86ad48","from":{"address":"7f56073741f18d4796870132a1107087d11c5e7e","owner":"unknown","owner_type":"unknown"},"to":{"address":"7122db0ebe4eb9b434a9f2ffe6760bc03bfbd0e0","owner":"unknown","owner_type":"unknown"},"timestamp":1679758751,"amount":824064.6,"amount_usd":830320.06,"transaction_count":1},{"blockchain":"ethereum","symbol":"usdc","id":"1990512548","transaction_type":"transfer","hash":"b13a7ba1d0232779fa8465715a5401a7b145271a1146415d34f34ee2dc86ad48","from":{"address":"c9f93163c99695c6526b799ebca2207fdf7d61ad","owner":"unknown","owner_type":"unknown"},"to":{"address":"7122db0ebe4eb9b434a9f2ffe6760bc03bfbd0e0","owner":"unknown","owner_type":"unknown"},"timestamp":1679758751,"amount":824930.4,"amount_usd":827961.25,"transaction_count":1}]}`))
			return
		}
		// When no transactions in given hash
		if r.Header.Get("X-WA-API-KEY") == "CORRECT" && strings.Contains(r.URL.Path, btcChain) && strings.Contains(r.URL.Path, hashCorrect) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result":"success","count":0}`))
			return
		}
		// When wrong hash length
		if r.Header.Get("X-WA-API-KEY") == "CORRECT" && strings.Contains(r.URL.Path, ethChain) && strings.Contains(r.URL.Path, hashIncorrect) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"result":"error","message":"invalid value for hash parameter"}`))
			return
		}
		// When wrong blockchain param
		if r.Header.Get("X-WA-API-KEY") == "CORRECT" && strings.Contains(r.URL.Path, incorrectChain) && strings.Contains(r.URL.Path, hashCorrect) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"result":"error","message":"invalid value for blockchain parameter"}`))
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"result":"error","message":"invalid api_key"}`))
	}))

	api := whalealertapi.New().WithCustomURL(serverOK.URL).WithAccessKey("CORRECT")
	res, err := api.Transaction(ethChain, hashCorrect)
	if err != nil {
		t.Errorf("Expected OK got error: %s", err)
	}
	if res.Result != "success" {
		t.Errorf("Expected %s got: %s", "success", res.Result)
	}

	res, err = api.Transaction(btcChain, hashCorrect)
	if err != nil {
		t.Errorf("Expected OK got error: %s", err)
	}
	if res.Result != "success" {
		t.Errorf("Expected %s got: %s", "success", res.Result)
	}
	if res.Count != 0 {
		t.Errorf("Expected %d got: %d", 0, res.Count)
	}

	res, err = api.Transaction(ethChain, hashIncorrect)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if res != nil {
		t.Errorf("Expected nil, got: %v", res)
	}
	if err.Error() != "invalid value for hash parameter" {
		t.Errorf("Expected %s got: %s", "invalid value for hash parameter", err)
	}

	res, err = api.Transaction(incorrectChain, hashCorrect)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if res != nil {
		t.Errorf("Expected nil, got: %v", res)
	}
	if err.Error() != "invalid value for blockchain parameter" {
		t.Errorf("Expected %s got: %s", "invalid value for hash parameter", err)
	}

}

const (
	startTimeEmpty       = "1679774558"
	startTimeWithRecords = "1679774508"
)

func TestTransactions(t *testing.T) {
	serverOK := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//When correct blockchain and hash
		if r.Header.Get("X-WA-API-KEY") == "CORRECT" && r.URL.Query().Get("start") == startTimeEmpty {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result":"success","cursor":"0-0-0","count":0}`))
			return
		}
		if r.Header.Get("X-WA-API-KEY") == "CORRECT" && r.URL.Query().Get("start") == startTimeWithRecords {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result":"success","cursor":"76a63333-76a63333-641f534f","count":2,"transactions":[{"blockchain":"ethereum","symbol":"usdc","id":"1990603462","transaction_type":"transfer","hash":"7fc67143eb1a43ea111921076a44f7e2fa79ffe530bd3f79dee0464cf6214ae5","from":{"address":"6224b133adbd85dddc03a8ae111dec912f0ee4a7","owner":"Kraken","owner_type":"exchange"},"to":{"address":"ae2d4617c862309a3d75a0ffb358c7a5009c673f","owner":"Kraken","owner_type":"exchange"},"timestamp":1679774519,"amount":5000000,"amount_usd":5050054.5,"transaction_count":1},{"blockchain":"ethereum","symbol":"usdc","id":"1990603571","transaction_type":"transfer","hash":"9873bb41c78396a68d9883f0b0c79fe1c83ab17a93829d074f4ca3cfaf416ceb","from":{"address":"88e6a0c2ddd26feeb64f039a2c41296fcb3f5640","owner":"unknown","owner_type":"unknown"},"to":{"address":"f8b721bff6bf7095a0e10791ce8f998baa254fd0","owner":"unknown","owner_type":"unknown"},"timestamp":1679774543,"amount":517499.72,"amount_usd":522680.34,"transaction_count":1}]}`))
			return
		}
	}))

	api := whalealertapi.New().WithCustomURL(serverOK.URL).WithAccessKey("CORRECT")
	res, err := api.Transactions(1679774558, whalealertapi.TransactionsRequest{})
	if err != nil {
		t.Errorf("Expected OK got error: %s", err)
	}
	if res.Result != "success" {
		t.Errorf("Expected %s got: %s", "success", res.Result)
	}
	if res.Cursor != "0-0-0" {
		t.Errorf("Expected %s got: %s", "0-0-0", res.Cursor)
	}
	if res.Count != 0 {
		t.Errorf("Expected %d got: %d", 0, res.Count)
	}

	res, err = api.Transactions(1679774508, whalealertapi.TransactionsRequest{})
	if err != nil {
		t.Errorf("Expected OK got error: %s", err)
	}
	if res.Result != "success" {
		t.Errorf("Expected %s got: %s", "success", res.Result)
	}
	if res.Cursor != "76a63333-76a63333-641f534f" {
		t.Errorf("Expected %s got: %s", "76a63333-76a63333-641f534f", res.Cursor)
	}
	if res.Count != 2 {
		t.Errorf("Expected %d got: %d", 2, res.Count)
	}
}
