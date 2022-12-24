package api

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sheacloud/tfom/internal/identifiers"
	"github.com/sheacloud/tfom/pkg/execution/models"
)

func (c *ExecutionAPIClient) GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error) {
	applyExecutionRequest, err := c.dbClient.GetApplyExecutionRequest(ctx, applyExecutionRequestId)
	if err != nil {
		return nil, err
	}

	// fetch init and apply outputs from S3
	if applyExecutionRequest.InitOutputKey != "" {
		initOutput, err := c.DownloadTerraformApplyInitResults(ctx, applyExecutionRequest.InitOutputKey)
		if err != nil {
			return nil, err
		}
		applyExecutionRequest.InitOutput = initOutput
	}

	if applyExecutionRequest.ApplyOutputKey != "" {
		applyOutput, err := c.DownloadTerraformApplyResults(ctx, applyExecutionRequest.ApplyOutputKey)
		if err != nil {
			return nil, err
		}
		applyExecutionRequest.ApplyOutput = applyOutput
	}

	return applyExecutionRequest, nil
}

func (c *ExecutionAPIClient) GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	// TODO: fetch outputs from S3
	return c.dbClient.GetApplyExecutionRequests(ctx, limit, cursor)
}

func (c *ExecutionAPIClient) GetApplyExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	// TODO: fetch outputs from S3
	return c.dbClient.GetApplyExecutionRequestsByStateKey(ctx, stateKey, limit, cursor)
}

func (c *ExecutionAPIClient) GetApplyExecutionRequestsByGroupingKey(ctx context.Context, groupingKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	// TODO: fetch outputs from S3
	return c.dbClient.GetApplyExecutionRequestsByGroupingKey(ctx, groupingKey, limit, cursor)
}

func (c *ExecutionAPIClient) PutApplyExecutionRequest(ctx context.Context, input *models.NewApplyExecutionRequest) (*models.ApplyExecutionRequest, error) {
	applyExecutionRequestId, err := identifiers.NewIdentifier(identifiers.ResourceTypeApplyExecutionRequest)
	if err != nil {
		return nil, err
	}

	workflowExecutionId := uuid.New().String()

	applyExecutionRequest := models.ApplyExecutionRequest{
		ApplyExecutionRequestId:      applyExecutionRequestId.String(),
		TerraformVersion:             input.TerraformVersion,
		StateKey:                     input.StateKey,
		GroupingKey:                  input.GroupingKey,
		TerraformConfigurationBase64: input.TerraformConfigurationBase64,
		TerraformPlanBase64:          input.TerraformPlanBase64,
		AdditionalArguments:          input.AdditionalArguments,
		WorkflowExecutionId:          workflowExecutionId,
		Status:                       models.ApplyExecutionStatusPending,
		RequestTime:                  time.Now().UTC(),
	}
	err = c.dbClient.PutApplyExecutionRequest(ctx, &applyExecutionRequest)
	if err != nil {
		return nil, err
	}

	// Start Workflow
	err = c.startTerraformExecutionWorkflow(ctx, &TerraformExecutionWorkflowInput{
		RequestType: "apply",
		RequestId:   applyExecutionRequestId.String(),
	})
	if err != nil {
		return nil, err
	}

	return &applyExecutionRequest, nil
}

func (c *ExecutionAPIClient) UpdateApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, input *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error) {
	return c.dbClient.UpdateApplyExecutionRequest(ctx, applyExecutionRequestId, input)
}

func (c *ExecutionAPIClient) UploadTerraformApplyInitResults(ctx context.Context, planExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error) {
	return c.dbClient.UploadTerraformApplyInitResults(ctx, planExecutionRequestId, initResults)
}

func (c *ExecutionAPIClient) UploadTerraformApplyResults(ctx context.Context, planExecutionRequestId string, planResults *models.TerraformApplyOutput) (string, error) {
	return c.dbClient.UploadTerraformApplyResults(ctx, planExecutionRequestId, planResults)
}

func (c *ExecutionAPIClient) DownloadTerraformApplyInitResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformInitOutput, error) {
	return c.dbClient.DownloadTerraformApplyInitResults(ctx, applyExecutionRequestId)
}

func (c *ExecutionAPIClient) DownloadTerraformApplyResults(ctx context.Context, applyExecutionRequestId string) (*models.TerraformApplyOutput, error) {
	return c.dbClient.DownloadTerraformApplyResults(ctx, applyExecutionRequestId)
}
