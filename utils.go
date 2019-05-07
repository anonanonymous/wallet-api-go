package walletapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// makeJSONString - converts a map[string]interface to a string
func makeJSONString(dict map[string]interface{}) string {
	result, err := json.Marshal(dict)
	if err != nil {
		panic(err)
	}
	return string(result)
}

// sendRequest - helper function for sending http requests to wallet-api
func (wAPI WalletAPI) sendRequest(method, uri, data string) (*map[string]interface{}, *[]byte, error) {
	var rawData []byte
	var body map[string]interface{}

	req, err := http.NewRequest(method, uri, bytes.NewBufferString(data))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("X-API-KEY", wAPI.APIKey)
	req.Header.Add("Content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	req.Close = true
	if resp.StatusCode > 400 {
		return nil, nil, errors.New(ERRORS[resp.StatusCode])
	}

	defer resp.Body.Close()

	rawData, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		json.Unmarshal(rawData, &body)
	}

	if resp.StatusCode == 400 {
		return nil, nil, errors.New(body["errorMessage"].(string))
	}

	return &body, &rawData, err
}
