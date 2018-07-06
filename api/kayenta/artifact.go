/*
 * Kayenta API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kayenta

type Artifact struct {

	ArtifactAccount string `json:"artifactAccount,omitempty"`

	Location string `json:"location,omitempty"`

	Metadata *interface{} `json:"metadata,omitempty"`

	Name string `json:"name,omitempty"`

	Provenance string `json:"provenance,omitempty"`

	Reference string `json:"reference,omitempty"`

	Type_ string `json:"type,omitempty"`

	Uuid string `json:"uuid,omitempty"`

	Version string `json:"version,omitempty"`
}
