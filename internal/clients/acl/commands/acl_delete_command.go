package commands

import (
	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// Errors
const (
	errOnlyOneTopicOrConsumerGroupAllowed = "At most one of either a topic or consumer group can be specified"
)

// ACLDeleteCommand is a struct for ACL delete command
type ACLDeleteCommand exec.Cmd

// NewACLDeleteCommand is a factory method for ACL delete command
func NewACLDeleteCommand(acl v1alpha1.ACLBlock) (ACLDeleteCommand, error) {
	var command = ACLDeleteCommand{
		Path: clients.CliName,
		Args: []string{"kafka", "acl", "delete", "--environment", acl.Environment, "--cluster", acl.Cluster, "-o", "json"},
	}

	command.Args = append(command.Args, "--operation", acl.Operation)

	// Do some cast/assertion to reuse parseX funcs
	cmd := interface{}(command).(exec.Cmd)
	err := parsePatternType(&cmd, acl.PatternType)
	if err != nil {
		return command, err
	}

	err = parsePermission(&cmd, acl.Permission)
	if err != nil {
		return command, err
	}

	err = parseServiceAccount(&cmd, acl.Principal)
	if err != nil {
		return command, err
	}

	err = parseResource(&cmd, acl.ResourceName, acl.ResourceType)
	if err != nil {
		return command, nil
	}
	command.Args = append(command.Args, "--operation", acl.Operation)

	command = interface{}(command).(ACLDeleteCommand)

	return command, nil
}
