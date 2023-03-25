package whalealertapi

import (
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
)

func TestToURLArguments(t *testing.T) {
	args := []APIArgument{
		{
			Key:   "testKey1",
			Value: "testValue1",
		},
		{
			Key:   "testKey1",
			Value: "testValue1",
		},
	}
	expected := "testKey1=testValue1&testKey1=testValue1"
	actual := toURLArguments(args)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	args = []APIArgument{}
	expected = ""
	actual = toURLArguments(args)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}

type response struct {
	Response string `json:"response"`
}

func TestGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/ok") && r.Header.Get("X-WA-API-KEY") == "OK" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"response": "OK!"}`))
			return
		}
		if strings.Contains(r.URL.Path, "/bad_request_api") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "Bad Request - Result from api", "result": "error"}`))
			return
		}
		if strings.Contains(r.URL.Path, "/malformed_json") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Bad Request - malformed json", "result": "error"`))
			return
		}
		if strings.Contains(r.URL.Path, "/bad_request_mj") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "Bad Request - malformed json", "result": "error"`))
			return
		}
		if strings.Contains(r.URL.Path, "/wrong_json_format") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"abc": "error"}`))
			return
		}
	}))
	defer server.Close()

	client := &http.Client{}
	res, err := get[response](client, server.URL, "OK", "/ok", []APIArgument{})
	if err != nil {
		t.Errorf("Expected %s, got %s", "nil", err)
	}
	if res.Response != "OK!" {
		t.Errorf("Expected %s, got %s", "OK!", res.Response)
	}

	res, err = get[response](client, server.URL, "NOT_OK", "/bad_request_api", []APIArgument{})
	if err == nil {
		t.Errorf("Expected error")
	}
	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
	if err.Error() != "Bad Request - Result from api" {
		t.Errorf("Expected %s, got %s", "Bad Request", err)
	}

	res, err = get[response](client, server.URL, "NOT_OK", "/malformed_json", []APIArgument{})
	if err == nil {
		t.Errorf("Expected error")
	}
	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
	if err.Error() != "unexpected EOF" {
		t.Errorf("Expected %s, got %s", "unexpected EOF", err)
	}
	res, err = get[response](client, server.URL, "NOT_OK", "/bad_request_mj", []APIArgument{})
	if err == nil {
		t.Errorf("Expected error")
	}
	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
	if err.Error() != "unexpected EOF" {
		t.Errorf("Expected %s, got %s", "unexpected EOF", err)
	}
	res, err = get[response](client, server.URL, "NOT_OK", "/wrong_json_format", []APIArgument{})
	if err == nil {
		t.Errorf("Expected error")
	}
	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
	if err.Error() != "incorrect JSON response" {
		t.Errorf("Expected %s, got %s", "incorrect JSON response", err)
	}

	// Test required fields
	res, err = get[response](client, "", "NOT_OK", "/wrong_json_format", []APIArgument{})
	if err == nil {
		t.Errorf("Expected error")
	}
	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
	if err.Error() != "url is missing" {
		t.Errorf("Expected %s, got %s", "url is missing", err)
	}
	res, err = get[response](client, server.URL, "", "/wrong_json_format", []APIArgument{})
	if err == nil {
		t.Errorf("Expected error")
	}
	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
	if err.Error() != "access key is missing" {
		t.Errorf("Expected %s, got %s", "access key is missing", err)
	}
}

func TestRequiredFields(t *testing.T) {
	err := checkRequiredFields("", "abc")
	if err != ErrMissingURL {
		t.Errorf("Expected ErrMissingURL error")
	}
	err = checkRequiredFields("http://some_domain.com", "")
	if err != ErrMissingAccessKey {
		t.Errorf("Expected ErrMissingAccessKey error")

	}
}

func TestToAPIArguments(t *testing.T) {
	transaction := TransactionsRequest{
		Start:    1679760122,
		End:      1679760126,
		Limit:    10,
		Cursor:   "random",
		MinValue: 500000,
		Currency: "eth",
	}
	expected := []APIArgument{
		{Key: "start", Value: "1679760122"},
		{Key: "end", Value: "1679760126"},
		{Key: "limit", Value: "10"},
		{Key: "cursor", Value: "random"},
		{Key: "min_value", Value: "500000"},
		{Key: "currency", Value: "eth"},
	}
	sort.SliceStable(expected, func(i, j int) bool {
		return expected[i].Key < expected[j].Key
	})

	got := transaction.toAPIArguments()
	sort.SliceStable(got, func(i, j int) bool {
		return got[i].Key < got[j].Key
	})

	if len(got) != len(expected) {
		t.Errorf("Expected %d arguments, got %d", len(expected), len(got))
	}
	for i, arg := range got {
		if arg.Key != expected[i].Key || arg.Value != expected[i].Value {
			t.Errorf("Expected %s, got %s", expected[i].Key, arg.Key)
			t.Errorf("Expected %s, got %s", expected[i].Value, arg.Value)
		}
	}

	transaction1 := TransactionsRequest{
		Start: 1679760122,
		Limit: 1,
	}

	expected1 := []APIArgument{
		{Key: "start", Value: "1679760122"},
		{Key: "limit", Value: "1"},
	}
	sort.SliceStable(expected1, func(i, j int) bool {
		return expected1[i].Key < expected1[j].Key
	})

	got1 := transaction1.toAPIArguments()
	sort.SliceStable(got1, func(i, j int) bool {
		return got1[i].Key < got1[j].Key
	})

	if len(got1) != len(expected1) {
		t.Errorf("Expected %d arguments, got %d", len(expected1), len(got1))
	}
	for i, arg := range got1 {
		if arg.Key != expected1[i].Key || arg.Value != expected1[i].Value {
			t.Errorf("Expected %s, got %s", expected[i].Key, arg.Key)
			t.Errorf("Expected %s, got %s", expected[i].Value, arg.Value)
		}
	}
}
