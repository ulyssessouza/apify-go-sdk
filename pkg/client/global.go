package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const kvUrlFormat = "https://api.apify.com/v2/key-value-stores/%s/records/%s"

func GetInput() ([]byte, error) {
	actor := &Actor{}
	kvStoreId, ok := actor.GetDefaultKeyValueStoreId()
	if !ok {
		return nil, _E("could not get default key/value store ID")
	}
	inputKey, ok := actor.GetInputKey()
	if !ok {
		return nil, _E("could not get input key")
	}
	resp, err := http.Get(_S(kvUrlFormat, kvStoreId, inputKey))
	if err != nil {
		return nil, _E("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func SetOutput(output string) (string, error) {
	apify := &Apify{}
	actor := &Actor{}
	token, ok := apify.GetToken()
	if !ok {
		return "", _E("could not get token from environment")
	}
	kvStoreId, ok := actor.GetDefaultKeyValueStoreId()
	if !ok {
		return "", _E("could not get default key/value store ID")
	}
	url := _S(kvUrlFormat, kvStoreId, apify.GetOutputKvKey())
	dataMap := map[string]string{
		"out": output,
	}
	b, err := json.Marshal(dataMap)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", _S("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
