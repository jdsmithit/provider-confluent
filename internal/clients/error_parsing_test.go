// nolint
package clients

import (
	"testing"
)

type errorData struct {
	ErrorOutputFromCli string
	ErrorMessageString string
}

func TestErrorParsing1(t *testing.T) {
	errors := make(map[string]string)
	errors["xa"] = ErrorExampleApiKeyInvalidEnvironment
	errors["xa"] = ErrorExampleApiKeyInvalidEnvironment
	errors["xa"] = ErrorExampleApiKeyInvalidEnvironment
	errors["xa"] = ErrorExampleApiKeyInvalidEnvironment

	for _, v := range errors {
		ErrorParser([]byte(v))
	}
}
