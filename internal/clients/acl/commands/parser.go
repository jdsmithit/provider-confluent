package commands

import (
	"errors"
	"os/exec"
	"strings"
)

const (
	errPatternInvalid               = "pattern type must be either LITERAL or PREFIXED"
	errActionInvalid                = "action type must be either ALLOW or DENY"
	errPrincipalInvalid             = "principal does only allow User:sa-55555 type input"
	errResourceTypeInvalid          = "resource type must be either TOPIC, CONSUMER_GROUP or CLUSTER"
	errNoResourceNameForTypeCluster = "cannot use resource name when type is CLUSTER"
)

func parsePatternType(cmd *exec.Cmd, patternType string) error {
	switch patternType {
	case "PREFIXED":
		cmd.Args = append(cmd.Args, "--prefix")
		return nil
	case "LITERAL":
		return nil
	default:
		return errors.New(errPatternInvalid)
	}
}

func parseAction(cmd *exec.Cmd, action string) error {
	switch action {
	case "ALLOW":
		cmd.Args = append(cmd.Args, "--allow")
		return nil
	case "DENY":
		cmd.Args = append(cmd.Args, "--deny")
		return nil
	default:
		return errors.New(errActionInvalid)
	}
}

func parseServiceAccount(cmd *exec.Cmd, principal string) error {
	split := strings.Split(principal, ":")
	if len(split) != 2 {
		return errors.New(errPrincipalInvalid)
	}
	user := split[0]
	if user != "User" {
		return errors.New(errPrincipalInvalid)
	}

	serviceaccount := split[1]
	if strings.Contains(serviceaccount, "sa-") {
		return errors.New(errPrincipalInvalid)
	}

	cmd.Args = append(cmd.Args, "--service-account", serviceaccount)
	return nil
}

func parseResource(cmd *exec.Cmd, rName string, rType string) error {
	switch rType {
	case "TOPIC":
		cmd.Args = append(cmd.Args, "--topic", rName)
		return nil
	case "CONSUMER_GROUP":
		cmd.Args = append(cmd.Args, "--consumer-group", rName)
		return nil
	case "CLUSTER":
		if rName != "" {
			return errors.New(errNoResourceNameForTypeCluster)
		}
		cmd.Args = append(cmd.Args, "--cluster-scope")
		return nil
	default:
		return errors.New(errResourceTypeInvalid)
	}
}
