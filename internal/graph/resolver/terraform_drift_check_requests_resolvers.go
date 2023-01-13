package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/tfom/internal/graph/generated"
	"github.com/sheacloud/tfom/pkg/models"
)

// CreateTerraformDriftCheckRequest is the resolver for the createTerraformDriftCheckRequest field.
func (r *mutationResolver) CreateTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequest models.NewTerraformDriftCheckRequest) (*models.TerraformDriftCheckRequest, error) {
	return r.apiClient.PutTerraformDriftCheckRequest(ctx, &terraformDriftCheckRequest)
}

// ModuleAssignment is the resolver for the moduleAssignment field.
func (r *terraformDriftCheckRequestResolver) ModuleAssignment(ctx context.Context, obj *models.TerraformDriftCheckRequest) (*models.ModuleAssignment, error) {
	return r.apiClient.GetModuleAssignment(ctx, obj.ModuleAssignmentId)
}

// PlanExecutionRequest is the resolver for the planExecutionRequest field.
func (r *terraformDriftCheckRequestResolver) PlanExecutionRequest(ctx context.Context, obj *models.TerraformDriftCheckRequest, withOutputs *bool) (*models.PlanExecutionRequest, error) {
	if obj.PlanExecutionRequestId == nil {
		return nil, nil
	}
	return r.apiClient.GetPlanExecutionRequest(ctx, *obj.PlanExecutionRequestId, aws.ToBool(withOutputs))
}

// ModulePropagation is the resolver for the modulePropagation field.
func (r *terraformDriftCheckRequestResolver) ModulePropagation(ctx context.Context, obj *models.TerraformDriftCheckRequest) (*models.ModulePropagation, error) {
	if obj.ModulePropagationId == nil {
		return nil, nil
	}
	return r.apiClient.GetModulePropagation(ctx, *obj.ModulePropagationId)
}

// ModulePropagationDriftCheckRequest is the resolver for the modulePropagationDriftCheckRequest field.
func (r *terraformDriftCheckRequestResolver) ModulePropagationDriftCheckRequest(ctx context.Context, obj *models.TerraformDriftCheckRequest) (*models.ModulePropagationDriftCheckRequest, error) {
	if obj.ModulePropagationId == nil || obj.ModulePropagationDriftCheckRequestId == nil {
		return nil, nil
	}
	return r.apiClient.GetModulePropagationDriftCheckRequest(ctx, *obj.ModulePropagationId, *obj.ModulePropagationDriftCheckRequestId)
}

// TerraformDriftCheckRequest returns generated.TerraformDriftCheckRequestResolver implementation.
func (r *Resolver) TerraformDriftCheckRequest() generated.TerraformDriftCheckRequestResolver {
	return &terraformDriftCheckRequestResolver{r}
}

type terraformDriftCheckRequestResolver struct{ *Resolver }