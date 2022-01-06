package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	ErrNotFound             = "schema not found"
	ErrNotParsed            = "cannot be parsed"
	ErrGeneral              = "unknown error:"
	ErrInvalidResponse      = "invalid response from describe"
	ErrNotCompatible        = "schema not compatible"
	ErrInvalidCompatibility = "invalid compatibility level"
	ErrUnknownFormat        = "unknown error format"
)

type ConfluentError struct {
	ErrorCode int64

	Err error
}

func NewConfluentError(err error) *ConfluentError {
	payload := &ConfluentError{
		ErrorCode: 0,
		Err:       err,
	}
	return payload
}

func (c *ConfluentError) Error() string {
	currentTimestamp := time.Now()
	msg := fmt.Sprintf("%s: ConfluentClient error || %s", currentTimestamp, c.Err.Error())

	return msg
}

type errorResponse struct {
	ErrorCode int64  `json:"error_code"`
	Message   string `json:"message"`
}

func ResponseSanitiser(cmdoutput []byte) ([]string, error) {
	out := string(cmdoutput)
	split := strings.SplitN(out, ":", 2)
	if len(split) > 2 {
		return []string{}, NewConfluentError(errors.New("ErrInvalidResponse"))
	}
	return split, nil
}

func WrapError(err error, newErr error) error {
	newMsg := fmt.Sprintf("%s\n%s", err.Error(), newErr.Error())
	return errors.New(newMsg)
}

func ErrorParser(cmdout []byte) error {
	split, err := ResponseSanitiser(cmdout)
	if err != nil {
		return err
	}

	var schema errorResponse

	if len(split) <= 1 {
		return NewConfluentError(errors.New(ErrUnknownFormat))
	}

	fmt.Println(split[1])

	err = json.Unmarshal([]byte(split[1]), &schema)
	if err != nil {
		message := fmt.Sprintf("%s%v", ErrGeneral, split)
		return NewConfluentError(WrapError(err, errors.New(message)))
	}

	switch schema.ErrorCode {
	case 409:
		return NewConfluentError(errors.New(ErrNotCompatible))
	case 40401:
		return NewConfluentError(errors.New(ErrNotFound))
	case 42203:
		return NewConfluentError(errors.New(ErrInvalidCompatibility))
	default:
		return NewConfluentError(errors.New(fmt.Sprintf("%s%d%s", ErrGeneral, schema.ErrorCode, schema.Message)))
	}
}
