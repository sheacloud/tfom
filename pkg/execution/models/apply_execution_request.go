package models

import "time"

type ApplyExecutionStatus string

const (
	ApplyExecutionStatusPending   ApplyExecutionStatus = "PENDING"
	ApplyExecutionStatusRunning   ApplyExecutionStatus = "RUNNING"
	ApplyExecutionStatusSucceeded ApplyExecutionStatus = "SUCCEEDED"
	ApplyExecutionStatusFailed    ApplyExecutionStatus = "FAILED"
)

type ApplyExecutionRequest struct {
	ApplyExecutionRequestId      string                `json:"applyExecutionRequestId"`
	TerraformVersion             string                `json:"terraformVersion"`
	StateKey                     string                `json:"stateKey"`
	TerraformConfigurationBase64 string                `json:"terraformConfigurationBase64"`
	TerraformPlanBase64          string                `json:"terraformPlanBase64"`
	AdditionalArguments          []string              `json:"additionalArguments"`
	WorkflowExecutionId          string                `json:"-"`
	Status                       ApplyExecutionStatus  `json:"status"`
	InitOutputKey                string                `json:"-"`                                   // for internal use only
	InitOutput                   *TerraformInitOutput  `json:"initOutput,omitempty" dynamodbav:"-"` // fetched from S3 on request
	ApplyOutputKey               string                `json:"-"`                                   // for internal use only
	ApplyOutput                  *TerraformApplyOutput `json:"planOutput,omitempty" dynamodbav:"-"` // fetched from S3 on request
	RequestTime                  time.Time             `json:"requestTime"`
}

type ApplyExecutionRequests struct {
	Items      []ApplyExecutionRequest `json:"items"`
	NextCursor string                  `json:"nextCursor"`
}

type NewApplyExecutionRequest struct {
	TerraformVersion             string   `json:"terraformVersion"`
	StateKey                     string   `json:"stateKey"`
	TerraformConfigurationBase64 string   `json:"terraformConfigurationBase64"`
	TerraformPlanBase64          string   `json:"terraformPlanBase64"`
	AdditionalArguments          []string `json:"additionalArguments"`
}

type ApplyExecutionRequestUpdate struct {
	InitOutputKey  *string
	ApplyOutputKey *string
	Status         *ApplyExecutionStatus
}