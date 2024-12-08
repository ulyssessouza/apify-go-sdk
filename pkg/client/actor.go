package client

import (
	"os"
)

type Actor struct{}

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
