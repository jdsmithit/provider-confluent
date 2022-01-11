package acl

import (
	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// IClient interface for service account client
type IClient interface {
	ACLCreate(acl v1alpha1.ACLBlock) (ACLResponseBlock, error)
	ACLDelete(acl v1alpha1.ACLBlock) error
	ACLList(serviceAccount string, environment string, cluster string) (ACLResponseBlockList, error)
	// ACLUpdate() error
}

// Config is a configuration element for the service account client
type Config struct {
	APICredentials clients.APICredentials
}

// Client is a struct for service account client
type Client struct {
	Config Config
}

type ACLResponseBlock struct {
	Operation    string `json:"operation"`
	PatternType  string `json:"pattern_type"`
	Permission   string `json:"permission"`
	Principal    string `json:"principal"`
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

type ACLResponseBlockList []ACLResponseBlock
