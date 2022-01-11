package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// NewACLCreateCommand is a factory method for ACL create command
func NewACLCreateCommand(acl v1alpha1.ACLBlock) (exec.Cmd, error) {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"kafka", "acl", "create", "--environment", acl.Environment, "--cluster", acl.Cluster, "--service-account", acl.ServiceAccount, "-o", "json"},
	}

	err := ACLCommonCommandParsing(&command, acl)
	if err != nil {
		return command, err
	}

	return command, nil
}

func ACLCommonCommandParsing(cmd *exec.Cmd, acl v1alpha1.ACLBlock) error {
	for _, v := range acl.Operations {
		cmd.Args = append(cmd.Args, "--operation", v)
	}

	err := parseAction(cmd, acl.Action)
	if err != nil {
		return err
	}

	if acl.Prefix {
		cmd.Args = append(cmd.Args, "--prefix")
	}

	err = parseResource(cmd, acl.ResourceName, acl.ResourceType)
	if err != nil {
		return err
	}

	return nil
}
