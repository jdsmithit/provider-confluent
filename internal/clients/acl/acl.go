package acl

import (
	"encoding/json"
	"os/exec"

	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"

	"github.com/pkg/errors"

	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/dfds/provider-confluent/internal/clients/acl/commands"
)

// Errors
const (
	errUnknown = "unknown error"
	// errEnvironmentNotFound                  = "environment not found"
	// errServiceAccountNotFoundOrLimitReached = "service not found or limit reached"
	// errResourceNotFoundOrAccessForbidden    = "resource not found or access forbidden"
	ErrACLNotExistsOrInvalidServiceAccount = "acl for service account does not exists or invalid service account"
	// errUnknownApiKey                        = "unknow apikey"
)

// NewClient is a factory method for apikey client
func NewClient(c Config) IClient {
	return &Client{Config: c}
}

func (c *Client) ACLCreate(aclP v1alpha1.ACLParameters) ([]v1alpha1.ACLRule, error) {
	var resp []v1alpha1.ACLRule

	cmd, err := commands.NewACLCreateCommand(aclP)
	if err != nil {
		return resp, err
	}

	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return resp, errorParser(out)
	}

	var aclBlocks []ACLBlock
	err = json.Unmarshal(out, &aclBlocks)
	if err != nil {
		return resp, err
	}

	for _, block := range aclBlocks {
		resp = append(resp, FromACLBlockToACLRule(block))
	}

	return resp, nil
}

func (c *Client) ACLDelete(aclP v1alpha1.ACLParameters) error {
	cmd, err := commands.NewACLDeleteCommand(aclP)
	if err != nil {
		return err
	}

	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return errorParser(out)
	}

	return nil
}

func (c *Client) ACLList(serviceAccount string, environment string, cluster string) ([]v1alpha1.ACLRule, error) {
	var resp []v1alpha1.ACLRule

	cmd := commands.NewACLListCommand(environment, cluster, serviceAccount)
	out, err := clients.ExecuteCommand(exec.Cmd(cmd))

	if err != nil {
		return resp, errorParser(out)
	}

	var aclBlocks []ACLBlock
	err = json.Unmarshal(out, &aclBlocks)
	if err != nil {
		return resp, err
	}

	for _, block := range aclBlocks {
		resp = append(resp, FromACLBlockToACLRule(block))
	}

	if len(resp) == 0 {
		return resp, errors.New(ErrACLNotExistsOrInvalidServiceAccount)
	}

	return resp, nil
}

func errorParser(cmdout []byte) error {
	str := string(cmdout)
	// if strings.Contains(str, "Error: environment") && strings.Contains(str, "not found") {
	// 	return errors.New(errEnvironmentNotFound)
	// } else if strings.Contains(str, "Your Api Keys per User is currently limited to 10") {
	// 	return errors.New(errServiceAccountNotFoundOrLimitReached)
	// } else if strings.Contains(str, "Error: Kafka cluster not found or access forbidden") {
	// 	return errors.New(errResourceNotFoundOrAccessForbidden)
	// } else if strings.Contains(str, "Error: Unknown API key") {
	// 	return errors.New(errUnknownApiKey)
	// }
	return errors.Wrap(errors.New(errUnknown), string(str))
}
