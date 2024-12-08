package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Actor struct{}

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

func (a *Actor) GetBuildId() (string, bool) {
	return os.LookupEnv("ACTOR_BUILD_ID")
}

func (a *Actor) GetBuildNumber() (string, bool) {
	return os.LookupEnv("ACTOR_BUILD_NUMBER")
}

func (a *Actor) GetDefaultDatasetId() (string, bool) {
	return os.LookupEnv("ACTOR_DEFAULT_DATASET_ID")
}

func (a *Actor) GetDefaultKeyValueStoreId() (string, bool) {
	return os.LookupEnv("ACTOR_DEFAULT_KEY_VALUE_STORE_ID")
}

func (a *Actor) GetDefaultRequestQueueId() (string, bool) {
	return os.LookupEnv("ACTOR_DEFAULT_REQUEST_QUEUE_ID")
}

func (a *Actor) GetEventsWebSocketUrl() (string, bool) {
	return os.LookupEnv("ACTOR_EVENTS_WEBSOCKET_URL")
}

func (a *Actor) GetId() (string, bool) {
	return os.LookupEnv("ACTOR_ID")
}

func (a *Actor) GetInputKey() (string, bool) {
	return os.LookupEnv("ACTOR_INPUT_KEY")
}

func (a *Actor) GetMaxPaidDatasetItems() (string, bool) {
	return os.LookupEnv("ACTOR_MAX_PAID_DATASET_ITEMS")
}

func (a *Actor) GetMaxTotalChargeUSD() (string, bool) {
	return os.LookupEnv("ACTOR_MAX_TOTAL_CHARGE_USD")
}

func (a *Actor) GetMemoryMbytes() (string, bool) {
	return os.LookupEnv("ACTOR_MEMORY_MBYTES")
}

func (a *Actor) GetRunId() (string, bool) {
	return os.LookupEnv("ACTOR_RUN_ID")
}

func (a *Actor) GetStandbyPort() (string, bool) {
	return os.LookupEnv("ACTOR_STANDBY_PORT")
}

func (a *Actor) GetStandbyUrl() (string, bool) {
	return os.LookupEnv("ACTOR_STANDBY_URL")
}

func (a *Actor) GetStartedAt() (string, bool) {
	return os.LookupEnv("ACTOR_STARTED_AT")
}

func (a *Actor) GetTaskId() (string, bool) {
	return os.LookupEnv("ACTOR_TASK_ID")
}

func (a *Actor) GetTimeoutAt() (string, bool) {
	return os.LookupEnv("ACTOR_TIMEOUT_AT")
}

func (a *Actor) GetWebServerPort() (string, bool) {
	return os.LookupEnv("ACTOR_WEB_SERVER_PORT")
}

func (a *Actor) GetWebServerUrl() (string, bool) {
	return os.LookupEnv("ACTOR_WEB_SERVER_URL")
}
