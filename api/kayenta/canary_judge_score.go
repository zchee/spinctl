/*
 * Kayenta API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kayenta

type CanaryJudgeScore struct {

	Classification string `json:"classification,omitempty"`

	ClassificationReason string `json:"classificationReason,omitempty"`

	Score float64 `json:"score,omitempty"`
}
