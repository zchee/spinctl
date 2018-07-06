/*
 * clouddriver
 *
 * Cloud read and write operations
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddriver

type ResponseEntity struct {

	Body *interface{} `json:"body,omitempty"`

	StatusCode string `json:"statusCode,omitempty"`

	StatusCodeValue int32 `json:"statusCodeValue,omitempty"`
}
