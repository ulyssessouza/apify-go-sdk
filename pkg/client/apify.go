package client

import (
	"os"
)

const kvUrlFormat = "https://api.apify.com/v2/key-value-stores/%s/records/%s"

type Apify struct{}

func (a *Apify) GetOutputKvKey() string {
	return "OUTPUT"
}

func (a *Apify) GetApiBaseUrl() (string, bool) {
	return os.LookupEnv("APIFY_API_BASE_URL")
}

func (a *Apify) GetApiPublicBaseUrl() (string, bool) {
	return os.LookupEnv("APIFY_API_PUBLIC_BASE_URL")
}

func (a *Apify) GetDedicatedCPUs() (string, bool) {
	return os.LookupEnv("APIFY_DEDICATED_CPUS")
}

func (a *Apify) GetDisableOutdatedWarning() (string, bool) {
	return os.LookupEnv("APIFY_DISABLE_OUTDATED_WARNING")
}

func (a *Apify) GetFact() (string, bool) {
	return os.LookupEnv("APIFY_FACT")
}

func (a *Apify) GetHeadless() (string, bool) {
	return os.LookupEnv("APIFY_HEADLESS")
}

func (a *Apify) GetIsAtHome() (string, bool) {
	return os.LookupEnv("APIFY_IS_AT_HOME")
}

func (a *Apify) GetMetaOrigin() (string, bool) {
	return os.LookupEnv("APIFY_META_ORIGIN")
}

func (a *Apify) GetProxyPassword() (string, bool) {
	return os.LookupEnv("APIFY_PROXY_PASSWORD")
}

func (a *Apify) GetProxyHostname() (string, bool) {
	return os.LookupEnv("APIFY_PROXY_HOSTNAME")
}

func (a *Apify) GetProxyStatusUrl() (string, bool) {
	return os.LookupEnv("APIFY_PROXY_STATUS_URL")
}

func (a *Apify) GetProxyPort() (string, bool) {
	return os.LookupEnv("APIFY_PROXY_PORT")
}

func (a *Apify) GetSdkLatestVersion() (string, bool) {
	return os.LookupEnv("APIFY_SDK_LATEST_VERSION")
}

func (a *Apify) GetToken() (string, bool) {
	return os.LookupEnv("APIFY_TOKEN")
}

func (a *Apify) GetUserId() (string, bool) {
	return os.LookupEnv("APIFY_USER_ID")
}

func (a *Apify) GetWorkflowKey() (string, bool) {
	return os.LookupEnv("APIFY_WORKFLOW_KEY")
}

func (a *Apify) GetActorBuildId() (string, bool) {
	return os.LookupEnv("APIFY_ACTOR_BUILD_ID")
}

func (a *Apify) GetActorBuildNumber() (string, bool) {
	return os.LookupEnv("APIFY_ACTOR_BUILD_NUMBER")
}

func (a *Apify) GetEventsWsUrl() (string, bool) {
	return os.LookupEnv("APIFY_ACTOR_EVENTS_WS_URL")
}

func (a *Apify) GetActorId() (string, bool) {
	return os.LookupEnv("APIFY_ACTOR_ID")
}

func (a *Apify) GetActorRunId() (string, bool) {
	return os.LookupEnv("APIFY_ACTOR_RUN_ID")
}

func (a *Apify) GetTaskId() (string, bool) {
	return os.LookupEnv("APIFY_ACTOR_TASK_ID")
}

func (a *Apify) GetContainerPort() (string, bool) {
	return os.LookupEnv("APIFY_CONTAINER_PORT")
}

func (a *Apify) GetContainerUrl() (string, bool) {
	return os.LookupEnv("APIFY_CONTAINER_URL")
}

func (a *Apify) GetDefaultDatasetId() (string, bool) {
	return os.LookupEnv("APIFY_DEFAULT_DATASET_ID")
}

func (a *Apify) GetDefaultKeyValueStoreId() (string, bool) {
	return os.LookupEnv("APIFY_DEFAULT_KEY_VALUE_STORE_ID")
}

func (a *Apify) GetDefaultRequestQueueId() (string, bool) {
	return os.LookupEnv("APIFY_DEFAULT_REQUEST_QUEUE_ID")
}

func (a *Apify) GetInputKey() (string, bool) {
	return os.LookupEnv("APIFY_INPUT_KEY")
}

func (a *Apify) GetMemoryMbytes() (string, bool) {
	return os.LookupEnv("APIFY_MEMORY_MBYTES")
}

func (a *Apify) GetStartedAt() (string, bool) {
	return os.LookupEnv("APIFY_STARTED_AT")
}

func (a *Apify) GetTimeoutAt() (string, bool) {
	return os.LookupEnv("APIFY_TIMEOUT_AT")
}

func (a *Apify) GetActId() (string, bool) {
	return os.LookupEnv("APIFY_ACT_ID")
}

func (a *Apify) GetActRunId() (string, bool) {
	return os.LookupEnv("APIFY_ACT_RUN_ID")
}
