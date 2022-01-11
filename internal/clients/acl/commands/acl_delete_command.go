package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"

	"github.com/dfds/provider-confluent/internal/clients"
)

// NewACLDeleteCommand is a factory method for ACL delete command
func NewACLDeleteCommand(acl v1alpha1.ACLBlock) (exec.Cmd, error) {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"kafka", "acl", "delete", "--environment", acl.Environment, "--cluster", acl.Cluster, "--service-account", acl.ServiceAccount},
	}

	err := ACLCommonCommandParsing(&command, acl)
	if err != nil {
		return command, err
	}

	return command, nil
}
