package clients

const ErrorExampleApiKeyInvalidEnvironment = `Error: environment "example" not found in context "login-confluent.cloud@dfds.com-https://confluent.cloud"`
const ErrorExampleApiKeyInvalidClusterOrAccessForbidden = `Error: Kafka cluster not found or access forbidden: error describing kafka cluster: Bad Request"`
const ErrorExampleApiKeyFailedToParseServiceAccount = `Error: failed to parse service account id: ensure service account id begins with "sa-"`
const ErrorExampleApiKeyInvalidServiceAccountOrApiKeyLimitReached = `Error: CCloud backend error: 1 error occurred:
	* error creating api key: Your Api Keys per User is currently limited to 10`
