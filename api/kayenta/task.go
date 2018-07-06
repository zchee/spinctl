/*
 * Kayenta API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kayenta

type Task struct {

	EndTime int64 `json:"endTime,omitempty"`

	Id string `json:"id,omitempty"`

	ImplementingClass string `json:"implementingClass,omitempty"`

	LoopEnd bool `json:"loopEnd,omitempty"`

	LoopStart bool `json:"loopStart,omitempty"`

	Name string `json:"name,omitempty"`

	StageEnd bool `json:"stageEnd,omitempty"`

	StageStart bool `json:"stageStart,omitempty"`

	StartTime int64 `json:"startTime,omitempty"`

	Status string `json:"status,omitempty"`
}
