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

// PlanExecutionRequests is the resolver for the planExecutionRequests field.
func (r *modulePropagationExecutionRequestResolver) PlanExecutionRequests(ctx context.Context, obj *models.ModulePropagationExecutionRequest, limit *int, nextCursor *string) (*models.PlanExecutionRequests, error) {
	if limit == nil {
		limit = aws.Int(100)
	}

	return r.apiClient.GetPlanExecutionRequestsByModulePropagationExecutionRequestId(ctx, obj.ModulePropagationExecutionRequestId, int32(*limit), aws.ToString(nextCursor))
}

// ApplyExecutionRequests is the resolver for the applyExecutionRequests field.
func (r *modulePropagationExecutionRequestResolver) ApplyExecutionRequests(ctx context.Context, obj *models.ModulePropagationExecutionRequest, limit *int, nextCursor *string) (*models.ApplyExecutionRequests, error) {
	if limit == nil {
		limit = aws.Int(100)
	}

	return r.apiClient.GetApplyExecutionRequestsByModulePropagationExecutionRequestId(ctx, obj.ModulePropagationExecutionRequestId, int32(*limit), aws.ToString(nextCursor))
}

// CreateModulePropagationExecutionRequest is the resolver for the createModulePropagationExecutionRequest field.
func (r *mutationResolver) CreateModulePropagationExecutionRequest(ctx context.Context, modulePropagationExecutionRequest models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error) {
	return r.apiClient.PutModulePropagationExecutionRequest(ctx, &modulePropagationExecutionRequest)
}

// ModulePropagationExecutionRequest is the resolver for the modulePropagationExecutionRequest field.
func (r *queryResolver) ModulePropagationExecutionRequest(ctx context.Context, modulePropagationID string, modulePropagationExecutionRequestID string) (*models.ModulePropagationExecutionRequest, error) {
	return r.apiClient.GetModulePropagationExecutionRequest(ctx, modulePropagationID, modulePropagationExecutionRequestID)
}

// ModulePropagationExecutionRequests is the resolver for the modulePropagationExecutionRequests field.
func (r *queryResolver) ModulePropagationExecutionRequests(ctx context.Context, limit *int, nextCursor *string) (*models.ModulePropagationExecutionRequests, error) {
	if limit == nil {
		limit = aws.Int(100)
	}

	return r.apiClient.GetModulePropagationExecutionRequests(ctx, int32(*limit), aws.ToString(nextCursor))
}

// ModulePropagationExecutionRequest returns generated.ModulePropagationExecutionRequestResolver implementation.
func (r *Resolver) ModulePropagationExecutionRequest() generated.ModulePropagationExecutionRequestResolver {
	return &modulePropagationExecutionRequestResolver{r}
}

type modulePropagationExecutionRequestResolver struct{ *Resolver }
