/*
 * Kayenta API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kayenta

type CanaryMetricConfig struct {

	AnalysisConfigurations map[string]map[string]interface{} `json:"analysisConfigurations,omitempty"`

	Groups []string `json:"groups,omitempty"`

	Name string `json:"name,omitempty"`

	Query *CanaryMetricSetQueryConfig `json:"query,omitempty"`

	ScopeName string `json:"scopeName,omitempty"`
}
