package whalealertapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var (
	ErrMissingURL       error = errors.New("url is missing")
	ErrMissingAccessKey error = errors.New("access key is missing")
	ErrIncorrectJSON    error = errors.New("incorrect JSON response")
	ErrNotFound         error = errors.New("endpoint not found")
)

// Single argument that can be passed to api endpoint
type APIArgument struct {
	Key   string `json:"name"`
	Value string `json:"value"`
}

// Converts Key and Value to url format key=value
func (arg APIArgument) toURLArgument() string {
	return fmt.Sprintf("%s=%v", arg.Key, arg.Value)
}

// Converts list of APIArgument to url format key1=value1&key2=value2
func toURLArguments(args []APIArgument) string {
	urlArgList := []string{}
	for _, arg := range args {
		urlArgList = append(urlArgList, arg.toURLArgument())
	}
	result := strings.Join(urlArgList, "&")
	return result
}

func isStructEmpty[T any](t T) bool {
	if !reflect.ValueOf(t).IsValid() {
		return true
	}
	if reflect.ValueOf(t).IsZero() {
		return true
	}
	return false
}

// get is doing get requests to specified url
// It returns T or error
func get[T any](client *http.Client, url string, key string, endpoint string, args []APIArgument) (*T, error) {
	err := checkRequiredFields(url, key)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s?%s", url, endpoint, toURLArguments(args)), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-WA-API-KEY", key)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		var result *T
		err = json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		if isStructEmpty(*result) {
			return nil, ErrIncorrectJSON
		}
		return result, nil

	}
	if response.StatusCode == 404 {
		return nil, ErrNotFound
	}
	var errResult *ErrorResponse
	err = json.NewDecoder(response.Body).Decode(&errResult)
	if err != nil {
		return nil, err
	}
	return nil, errResult
}

// checkRequiredFields returns error when url or key are empty
func checkRequiredFields(url, key string) error {
	if url == "" {
		return ErrMissingURL
	}
	if key == "" {
		return ErrMissingAccessKey
	}
	return nil
}

// toAPIArguments converts TransactionRequest to list of APIArguments
func (t TransactionsRequest) toAPIArguments() []APIArgument {
	args := []APIArgument{}
	v := reflect.ValueOf(t)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := typeOfS.Field(i)
		key := field.Tag.Get("arg")
		if key == "" {
			continue
		}
		value := v.Field(i).Interface()
		if reflect.ValueOf(value).IsValid() && !reflect.ValueOf(value).IsZero() {
			args = append(args, APIArgument{
				Key:   key,
				Value: fmt.Sprintf("%v", value),
			})
		}
	}
	return args
}
