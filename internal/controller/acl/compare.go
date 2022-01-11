package acl

import (
	"github.com/dfds/provider-confluent/apis/acl/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/acl"
)

func diffObserved(confluentResponse acl.ACLResponseBlockList, aclSpec []v1alpha1.ACLBlock) bool {

	return true
}

func diffEnvironment()

// type ACLBlock struct {
// 	Action string `json:"action"`
// 	// ClusterScope   string   `json:"clusterScope"`
// 	// ConsumerGroup  string   `json:"consumerGroup"`
// 	Operations     []string `json:"operations"`
// 	Prefix         bool     `json:"prefix"`
// 	ServiceAccount string   `json:"serviceAccount"`
// 	// Topic          string   `json:"topic"`
// 	Environment  string `json:"environment"`
// 	Cluster      string `json:"cluster"`
// 	ResourceType string `json:"resourceType"`
// 	ResourceName string `json:"resourceName"`
// }

// type ACLResponseBlock struct {
// 	Operation    string `json:"operation"`
// 	PatternType  string `json:"pattern_type"`
// 	Permission   string `json:"permission"`
// 	Principal    string `json:"principal"`
// 	ResourceName string `json:"resource_name"`
// 	ResourceType string `json:"resource_type"`
// }
