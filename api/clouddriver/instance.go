/*
 * clouddriver
 *
 * Cloud read and write operations
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddriver

type Instance struct {

	CloudProvider string `json:"cloudProvider,omitempty"`

	Health []Mapstringobject `json:"health,omitempty"`

	HealthState string `json:"healthState,omitempty"`

	LaunchTime int64 `json:"launchTime,omitempty"`

	Name string `json:"name,omitempty"`

	ProviderType string `json:"providerType,omitempty"`

	Zone string `json:"zone,omitempty"`
}
