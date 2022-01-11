package commands

import (
	"os/exec"

	"github.com/dfds/provider-confluent/internal/clients"
)

// NewACLListCommand is a factory method for ACL list command
func NewACLListCommand(serviceAccount string, environment string, cluster string) exec.Cmd {
	var command = exec.Cmd{
		Path: clients.CliName,
		Args: []string{"kafka", "acl", "list", "--environment", environment, "--cluster", cluster, "--service-account", serviceAccount, "-o", "json"},
	}

	return command
}
