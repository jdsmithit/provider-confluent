package schemaregistry

import "github.com/dfds/provider-confluent/internal/clients"

type Client interface {
	Create(subject string, schema string, schemaType string, environment string) (string, error)
	Delete(subject string, version string, permanent bool, environment string) (string, error)
	Describe(subject string, version string, environment string) (SchemaDescribeResponse, error)
}

type SchemaRegistryClient struct {
	Config clients.Config
}

type SchemaDescribeResponse struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Doc       string `json:"doc"`
	Fields    []struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Doc  string `json:"doc"`
	} `json:"fields"`
}