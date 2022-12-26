package workflow

import (
	"context"
	"encoding/base64"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskScheduleTerraformApply Task = "ScheduleTerraformApply"
)

type ScheduleTerraformApplyInput struct {
	Input struct {
		ModuleAccountAssociation            models.ModuleAccountAssociation
		ModulePropagationExecutionRequestId string
		ModulePropagationId                 string
	}
	TaskToken              string
	PlanExecutionRequestId string
}

type ScheduleTerraformApplyOutput struct {
	ApplyExecutionRequestId string
}

func (t *TaskHandler) ScheduleTerraformApply(ctx context.Context, input ScheduleTerraformApplyInput) (*ScheduleTerraformApplyOutput, error) {
	// get module propagation details
	modulePropagation, err := t.apiClient.GetModulePropagation(ctx, input.Input.ModulePropagationId)
	if err != nil {
		return nil, err
	}

	// get module version details
	moduleVersion, err := t.apiClient.GetModuleVersion(ctx, modulePropagation.ModuleGroupId, modulePropagation.ModuleVersionId)
	if err != nil {
		return nil, err
	}

	// get plan request details
	planRequest, err := t.apiClient.GetPlanExecutionRequest(ctx, input.PlanExecutionRequestId)
	if err != nil {
		return nil, err
	}

	// TODO: generate TF config file based on module version and account details

	planBase64 := base64.StdEncoding.EncodeToString(planRequest.PlanOutput.PlanFile)

	applyRequest, err := t.apiClient.PutApplyExecutionRequest(ctx, &models.NewApplyExecutionRequest{
		TerraformVersion:             moduleVersion.TerraformVersion,
		CallbackTaskToken:            input.TaskToken,
		StateKey:                     input.Input.ModuleAccountAssociation.RemoteStateKey,
		GroupingKey:                  input.Input.ModulePropagationExecutionRequestId,
		TerraformConfigurationBase64: "cHJvdmlkZXIgImF3cyIgewogIHJlZ2lvbiA9ICJ1cy1lYXN0LTEiCn0KCmRhdGEgImF3c19yZWdpb24iICJjdXJyZW50IiB7fQoKb3V0cHV0ICJyZWdpb24iIHsKICB2YWx1ZSA9IGRhdGEuYXdzX3JlZ2lvbi5jdXJyZW50Lm5hbWUKfQo=",
		TerraformPlanBase64:          planBase64,
	})
	if err != nil {
		return nil, err
	}

	return &ScheduleTerraformApplyOutput{
		ApplyExecutionRequestId: applyRequest.ApplyExecutionRequestId,
	}, nil
}
